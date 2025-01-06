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
	errUiMode                 = errors.New("no flags were provided")
)

type arguments struct {
	name     string
	path     string
	template string
}

// flags parses all flags and returns a structure with all possible arguments or
// and error that indicates to use the ui mode
func flags() (*arguments, error) {
	name := flag.String(
		"module",
		"",
		"your module name or path like github.com/you/proj",
	)

	path := flag.String(
		"path",
		"",
		"the path to put all the files",
	)

	template := flag.String(
		"template",
		"simple",
		"specify a template to avoid some boilerplate setup",
	)
	flag.Parse()

	args := &arguments{
		name:     *name,
		path:     *path,
		template: *template,
	}
	return args.validate()
}

// validate make sure that all arguments are set to create the project
func (args *arguments) validate() (*arguments, error) {
	if len(os.Args) < 2 {
		return nil, errUiMode
	}
	if len(args.name) == 0 {
		return nil, errModuleNameCannotBeZero
	}
	if len(args.path) == 0 {
		return nil, errPathCannotBeZero
	}

	return args, nil
}

func (args arguments) Name() string {
	return args.name
}

func (args arguments) Path() string {
	return args.path
}

func (args arguments) Template() string {
	return args.template
}

func ui() (*arguments, error) {
	return nil, nil
}

func run() error {
	args, err := flags()
	if shouldUseUiMode(err) {
		args, err = ui()
	}

	if err != nil {
		return err
	}

	if err := internal.CreateNewModule(args); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("could not create new project: %v\n", err)
		os.Exit(1)
	}
}

func shouldUseUiMode(err error) bool {
	return err != nil && err == errUiMode
}
