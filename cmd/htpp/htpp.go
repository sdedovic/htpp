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
	readFromStdin := flag.Bool("std-in", false, "Read JSON data from stdin")
	flag.Parse()

	var data map[string]any
	if *readFromStdin {
		err := json.NewDecoder(os.Stdin).Decode(&data)
		if err != nil {
			return err
		}
	}

	if filename := flag.Arg(0); filename != "" {
		template, err := htpp.Make(filename)
		if err != nil {
			return err
		}

		err = template.Execute(os.Stdout, data)
		if err != nil {
			return err
		}
	} else {
		return errors.New("no filename specified")
	}

	return nil
}

func main() {
	if err := mainErr(); err != nil {
		log.Fatalln("FATAL", fmt.Sprintf("main exited caused by: %v", err))
	}
}
