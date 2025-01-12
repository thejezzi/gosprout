package cli

import (
	"errors"
	"flag"
	"os"
)

var (
	errModuleNameCannotBeZero = errors.New("module name cannot be empty")
	errPathCannotBeZero       = errors.New("path cannot be empty")
	ErrUiMode                 = errors.New("no flags were provided")
)

type Arguments struct {
	name     string
	path     string
	template string
}

func NewArguments(moduleName, projectPath string) *Arguments {
	if len(projectPath) == 0 {
		projectPath = moduleName
	}
	return &Arguments{
		name:     moduleName,
		path:     projectPath,
		template: "simple",
	}
}

// flags parses all flags and returns a structure with all possible arguments or
// and error that indicates to use the ui mode
func Flags() (*Arguments, error) {
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

	// TODO: add templating when everything else is finished
	// template := flag.String(
	// 	"template",
	// 	"simple",
	// 	"specify a template to avoid some boilerplate setup",
	// )
	flag.Parse()

	return NewArguments(*name, *path).validate()
}

// validate make sure that all arguments are set to create the project
func (args *Arguments) validate() (*Arguments, error) {
	if len(os.Args) < 2 {
		return nil, ErrUiMode
	}
	if len(args.name) == 0 {
		return nil, errModuleNameCannotBeZero
	}
	if len(args.path) == 0 {
		return nil, errPathCannotBeZero
	}

	return args, nil
}

func (args Arguments) Name() string {
	return args.name
}

func (args Arguments) Path() string {
	return args.path
}

func (args Arguments) Template() string {
	return args.template
}

func (args Arguments) IsEmpty() bool {
	if len(args.name) == 0 {
		return false
	}
	if len(args.path) == 0 {
		return false
	}
	return true
}
