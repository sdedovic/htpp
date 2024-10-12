package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	htpp "github.com/sdedovic/htpp/internal"
	"log"
	"os"
)

func mainErr() error {
	readFromStdin := flag.Bool("std-in", false, "Read JSON applied to template from stdin")
	printDependencies := flag.Bool("print-dependencies", false, "Print all dependencies of the supplied template. Does not execute the template.")
	flag.Parse()

	filename := flag.Arg(0)
	if filename == "" {
		return errors.New("no filename specified")
	}

	var data map[string]any
	if *readFromStdin {
		err := json.NewDecoder(os.Stdin).Decode(&data)
		if err != nil {
			return err
		}
	}

	template, err := htpp.Make(filename)
	if err != nil {
		return err
	}

	if *printDependencies {
		for _, dep := range template.Dependencies {
			fmt.Println(dep)
		}
	} else {
		err = template.Inner.Execute(os.Stdout, data)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if err := mainErr(); err != nil {
		log.Fatalln("FATAL", fmt.Sprintf("main exited caused by: %v", err))
	}
}
