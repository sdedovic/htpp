package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"

	htpp "github.com/sdedovic/htpp/internal"
)

func run() error {
	readFromStdin := flag.Bool("stdin", false, "Read JSON applied to template from stdin")
	printDependencies := flag.Bool("print-dependencies", false, "Print all dependencies of the supplied template. Does not execute the template.")
	flag.Parse()

	filename := flag.Arg(0)
	if filename == "" {
		return errors.New("no filename specified")
	}

	var data map[string]any
	if *readFromStdin {
		if err := json.NewDecoder(os.Stdin).Decode(&data); err != nil {
			return err
		}
	}

	tmpl, err := htpp.Compile(filename)
	if err != nil {
		return err
	}

	if *printDependencies {
		for _, dep := range tmpl.Dependencies {
			fmt.Println(dep)
		}
		return nil
	}

	return tmpl.Execute(os.Stdout, data)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "htpp: %v\n", err)
		os.Exit(1)
	}
}
