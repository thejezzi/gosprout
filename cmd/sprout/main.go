package main

import (
	"fmt"
	"os"

	"github.com/thejezzi/gosprout/cmd/sprout/cli"
	"github.com/thejezzi/gosprout/internal/template"
	"github.com/thejezzi/gosprout/internal/ui"
	"github.com/thejezzi/gosprout/internal/util"
)

func run() error {
	var args *cli.Arguments
	var err error

	if len(os.Args) > 1 {
		args, err = cli.Flags()
	} else {
		args, err = ui.New()
	}

	if err != nil {
		return err
	}

	fmt.Println("creating project", args)
	for _, t := range template.All {
		if t.Name == args.Template() {
			return t.Create(args)
		}
	}

	return fmt.Errorf("template not found: %s", args.Template())
}

func main() {
	f, err := util.InitLogger()
	if err != nil {
		fmt.Printf("could not initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	if err := run(); err != nil {
		fmt.Printf("could not create new project: %v\n", err)
	}
}
