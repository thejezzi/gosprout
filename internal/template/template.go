package template

import (
	"fmt"
	"github.com/thejezzi/gosprout/cmd/sprout/cli"
	"github.com/thejezzi/gosprout/internal/git"
	"github.com/thejezzi/gosprout/internal/structure"
)

// Template struct and methods

type Template struct {
	Name        string
	description string
	Create      func(args *cli.Arguments) error
}

func New(name, description string, create func(args *cli.Arguments) error) Template {
	return Template{
		Name:        name,
		description: description,
		Create:      create,
	}
}

func (t Template) Title() string       { return t.Name }
func (t Template) Description() string { return t.description }
func (t Template) FilterValue() string { return t.Name }

// Template creation logic

func simpleCreate(args *cli.Arguments) error {
	return structure.CreateNewModule(args)
}

func testCreate(args *cli.Arguments) error {
	return structure.CreateNewModuleWithTest(args)
}

func gitCreate(args *cli.Arguments) error {
	if args.GitRepo == "" {
		return fmt.Errorf("git repository URL cannot be empty")
	}
	if err := git.Clone(args.GitRepo, args.Path()); err != nil {
		return err
	}
	if err := git.Reinit(args.Path()); err != nil {
		return err
	}
	return structure.ReplaceModuleName(args.Path(), args.Name())
}

// Template definitions

var (
	Simple = New(
		"Simple",
		"A simple structure with a cmd folder",
		simpleCreate,
	)
	Test = New(
		"Test",
		"A cmd folder with a main_test.go file",
		testCreate,
	)
	Git = New(
		"Git",
		"Create a project from a git repository",
		gitCreate,
	)
)

// All templates
var All = []Template{
	Simple,
	Test,
	Git,
}
