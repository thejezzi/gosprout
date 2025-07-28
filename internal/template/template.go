package template

import (
	"fmt"
	"os/exec"

	"github.com/thejezzi/mkgo/internal/args"
	"github.com/thejezzi/mkgo/internal/git"
)

// Template struct and methods

type Template struct {
	Name        string
	description string
	Create      func(args *args.Arguments) error
}

func New(
	name, description string,
	create func(args *args.Arguments) error,
) Template {
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

func simpleCreate(args *args.Arguments) error {
	err := CreateNewModule(args)
	if err != nil {
		return err
	}

	return nil
}

func testCreate(args *args.Arguments) error {
	if err := CreateNewModuleWithTest(args); err != nil {
		return err
	}

	if args.InitGit() {
		cmd := exec.Command("git", "init")
		cmd.Dir = args.Path()
		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	return nil
}

func gitCreate(args *args.Arguments) error {
	if args.GitRepo() == "" {
		return fmt.Errorf("git repository URL cannot be empty")
	}
	if err := git.Clone(args.GitRepo(), args.Path()); err != nil {
		return err
	}
	if err := git.Reinit(args.Path()); err != nil {
		return err
	}
	return ReplaceModuleName(args.Path(), args.Name())
}

// Template definitions

var (
	Simple = New(
		"Simple",
		"A simple template with a cmd folder",
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
