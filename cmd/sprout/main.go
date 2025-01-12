package main

import (
	"errors"
	"fmt"

	"github.com/thejezzi/gosprout/cmd/sprout/cli"
	"github.com/thejezzi/gosprout/internal/structure"
	"github.com/thejezzi/gosprout/internal/ui"
)

func runUI() (*cli.Arguments, error) {
	module, projectPath, err := ui.Run(
		ui.WithPrefixes(ui.FieldTitleModule, "github.com/thejezzi/"),
	)
	if err != nil {
		return nil, err
	}
	return cli.NewArguments(module, projectPath), nil
}

func run() error {
	args, err := cli.Flags()
	if errors.Is(err, cli.ErrUiMode) {
		err = nil
		args, err = runUI()
		if err != nil {
			return err
		}
	} else if err != nil {
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
