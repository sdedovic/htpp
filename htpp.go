package htpp

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	ErrInvalidExtends        = errors.New("extends has invalid syntax")
	ErrUnresolvableExtension = errors.New("could not resolve template extension")

	relativeFileRegex = regexp.MustCompile(`^\.{1,2}/`)
)

type Partial struct {
	identifier string
	Extends    string
	content    string
}

func (p Partial) String() string {
	return fmt.Sprintf("Partial{identifier: %q, extends: %q, content: [%db]}", p.identifier, p.Extends, len(p.content))
}

func parse(identifier string, src string) (Partial, error) {
	lines := strings.Split(src, "\n")

	if !strings.HasPrefix(lines[0], "extends") {
		return Partial{identifier: identifier, content: src}, nil
	}

	splits := strings.SplitN(lines[0], " ", 2)
	if len(splits) != 2 || splits[1] == "" {
		return Partial{}, ErrInvalidExtends
	}

	return Partial{
		identifier: identifier,
		Extends:    splits[1],
		content:    strings.Join(lines[1:], "\n"),
	}, nil
}

func parseFromFile(file string) (Partial, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return Partial{}, err
	}
	return parse(file, string(b))
}

// Template is a compiled htpp template ready for execution.
type Template struct {
	inner        *template.Template
	Dependencies []string
}

// Execute applies the template to data, writing the result to w.
func (t Template) Execute(w io.Writer, data any) error {
	return t.inner.Execute(w, data)
}

// Compile parses filename and resolves its extends chain into a Template.
func Compile(filename string) (Template, error) {
	stack := make([]*Partial, 0)
	dependencies := make([]string, 0)

	for {
		partial, err := parseFromFile(filename)
		if err != nil {
			return Template{}, err
		}

		stack = append(stack, &partial)

		if partial.Extends == "" {
			break
		}

		if relativeFileRegex.MatchString(partial.Extends) {
			filename = filepath.Clean(filepath.Join(filepath.Dir(filename), partial.Extends))
			dependencies = append(dependencies, filename)
			continue
		}

		return Template{}, ErrUnresolvableExtension
	}

	n := len(stack)
	tmpl, err := template.New("default").Parse(stack[n-1].content)
	if err != nil {
		return Template{}, err
	}

	stack = stack[:n-1]
	for len(stack) > 0 {
		n = len(stack)
		cloned, err := tmpl.Clone()
		if err != nil {
			return Template{}, err
		}
		tmpl, err = cloned.Parse(stack[n-1].content)
		if err != nil {
			return Template{}, err
		}
		stack = stack[:n-1]
	}

	return Template{tmpl, dependencies}, nil
}
