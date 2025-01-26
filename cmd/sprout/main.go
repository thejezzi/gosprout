package main

import (
	"fmt"
	"os"

	"github.com/thejezzi/gosprout/cmd/sprout/cli"
	"github.com/thejezzi/gosprout/internal/structure"
	"github.com/thejezzi/gosprout/internal/ui"
)

func runUI() (*cli.Arguments, error) {
	var module, projectPath string
	err := ui.Form(
		ui.Input().
			Title("module").
			Placeholder("you-awesome-module").
			Prompt("github.com/thejezzi/").
			Value(&module),
		ui.Input().
			Title("path").
			Placeholder("somewhere/to/put/project").
			Prompt("~/tmp/").
			FocusOnStart().
			Value(&projectPath),
		ui.List(),
	)
	if err != nil {
		return nil, err
	}

	return cli.NewArguments(module, projectPath), nil
}

func run() error {
	var args *cli.Arguments
	var err error

	if len(os.Args) > 1 {
		args, err = cli.Flags()
	} else {
		args, err = runUI()
	}

	if err != nil {
		return err
	}

	fmt.Println("creating project", args)
	if err := structure.CreateNewModule(args); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("could not create new project: %v\n", err)
	}
}
