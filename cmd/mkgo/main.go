package main

import (
	"fmt"
	"os"

	argsPkg "github.com/thejezzi/mkgo/internal/args"
	"github.com/thejezzi/mkgo/internal/template"
	"github.com/thejezzi/mkgo/internal/ui"
)

// newArgumentsFromUI converts a ui.UI to *argsPkg.Arguments
func newArgumentsFromUI(f ui.Form) *argsPkg.Arguments {
	// Converts a ui.Form to *argsPkg.Arguments
	if f == nil {
		return nil
	}
	return argsPkg.NewArguments(
		f.GetModule(),
		f.GetPath(),
		f.GetTemplate(),
		f.GetGitRepo(),
		f.GetCreateMakefile(),
		f.GetInitGit(),
	)
}

// getArguments returns Arguments from flags or UI
func getArguments() (*argsPkg.Arguments, error) {
	if len(os.Args) > 1 {
		return argsPkg.Flags()
	}
	form, err := ui.NewForm()
	if err != nil {
		return nil, err
	}
	return newArgumentsFromUI(form), nil
}

func run() error {
	args, err := getArguments()
	if err != nil {
		return err
	}

	for _, t := range template.All {
		if t.Name == args.Template() {
			err := t.Create(args)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return fmt.Errorf("template not found: %s", args.Template())
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("could not create new project: %v\n", err)
	}
}
