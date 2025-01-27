package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/thejezzi/gosprout/cmd/sprout/cli"
	"github.com/thejezzi/gosprout/internal/structure"
	"github.com/thejezzi/gosprout/internal/ui"
)

func runUI() (*cli.Arguments, error) {
	var module, projectPath, template string

	list := ui.List().SetItems(
		ui.ListItem("Simple", "A simple structure with a cmd folder"),
		ui.ListItem("Test", "A cmd folder with a main_test.go file"),
	)

	err := ui.Form(
		ui.Input().
			Title("module").
			Placeholder("you-awesome-module").
			Prompt("github.com/thejezzi/").
			Validate(func(s string) error {
				if len(s) == 0 {
					return errors.New("cannot be empty")
				}
				return nil
			}).
			Value(&module),
		ui.Input().
			Title("path").
			Placeholder("somewhere/to/put/project").
			Prompt("~/tmp/").
			FocusOnStart().
			Value(&projectPath),
		list.Title("template").
			Value(&template),
	)
	if err != nil {
		return nil, err
	}

	fmt.Println(template)
	return cli.NewArguments(module, projectPath, template), nil
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

func initLogger() *os.File {
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(f)
	return f
}

func main() {
	f := initLogger()
	defer f.Close()

	if err := run(); err != nil {
		fmt.Printf("could not create new project: %v\n", err)
	}
}
