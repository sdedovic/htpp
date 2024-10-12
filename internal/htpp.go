package internal

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"path"
	"regexp"
	"strings"
)

var (
	InvalidExtendsError              = errors.New("extends has invalid syntax")
	CouldNotResolveTemplateExtension = errors.New("could not resolve template extension")

	relativeFileRegex = regexp.MustCompile(`^\.{1,2}/`)
)

type Partial struct {
	identifier string
	Extends    string
	content    string
}

func (p Partial) String() string {
	return fmt.Sprintf("Partial{identifier: \"%s\", extends: \"%s\", content: [%db]}", p.identifier, p.Extends, len(p.content))
}

func parse(identifier string, template string) (Partial, error) {
	lines := strings.Split(template, "\n")

	if strings.HasPrefix(lines[0], "extends") {
		splits := strings.Split(lines[0], " ")
		if len(splits) != 2 {
			return Partial{}, InvalidExtendsError
		}

		return Partial{
			identifier: identifier,
			Extends:    splits[1],
			content:    strings.Join(lines[1:], "\n"),
		}, nil
	} else {
		return Partial{identifier: identifier, content: template}, nil
	}
}

func parseFromFile(file string) (Partial, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return Partial{}, err
	}
	return parse(file, string(b))
}

type Template struct {
	Inner        template.Template
	Dependencies []string
}

func Make(filename string) (Template, error) {
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
			filename = path.Clean(path.Join(path.Dir(filename), partial.Extends))
			dependencies = append(dependencies, filename)
			continue
		}

		return Template{}, CouldNotResolveTemplateExtension
	}

	n := len(stack)
	tmpl, err := template.New("default").Parse(stack[n-1].content)
	if err != nil {
		return Template{}, err
	}

	stack = stack[:n-1]
	for len(stack) > 0 {
		n := len(stack)
		tmpl, err = template.Must(tmpl.Clone()).Parse(stack[n-1].content)
		if err != nil {
			return Template{}, err
		}
		stack = stack[:n-1]
	}

	return Template{*tmpl, dependencies}, nil
}
