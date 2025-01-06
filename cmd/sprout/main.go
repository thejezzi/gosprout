package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/thejezzi/gosprout/internal"
)

var (
	errModuleNameCannotBeZero = errors.New("module name cannot be empty")
	errPathCannotBeZero       = errors.New("path cannot be empty")
)

type arguments struct {
	path     *string
	name     *string
	template *string
}

// flags parses all flags and returns a structure with all possible arguments or
// and error that indicates to use the ui mode
func flags() (*arguments, error) {
	args := &arguments{
		name: flag.String(
			"module",
			"",
			"your module name or path like github.com/you/proj",
		),

		path: flag.String(
			"path",
			"",
			"the path to put all the files",
		),

		template: flag.String(
			"template",
			"simple",
			"specify a template to avoid some boilerplate setup",
		),
	}
	flag.Parse()

	return args.validate()
}

// validate make sure that all arguments are set to create the project
func (args *arguments) validate() (*arguments, error) {
	if len(*args.name) == 0 {
		return nil, errModuleNameCannotBeZero
	}
	if len(*args.path) == 0 {
		return nil, errPathCannotBeZero
	}

	return args, nil
}

// exit prints the top level error and exits with status code one to indicate
// and error has happened
func exit(err error) {
	fmt.Printf("could not create new project: %v\n", err)
	os.Exit(1)
}

func main() {
	args, err := flags()
	if err != nil {
		exit(err)
	}

	if err := internal.CreateNewModule(*args.path, *args.name, *args.template); err != nil {
		exit(err)
	}
}
