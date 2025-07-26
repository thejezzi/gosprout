package ui

import (
	"errors"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/thejezzi/gosprout/cmd/sprout/cli"
)

func New() (*cli.Arguments, error) {
	var module, projectPath, template string

	fieldDefs := []FieldDef{
		{
			Title:       "module",
			Placeholder: "you-awesome-module",
			Prompt:      "github.com/thejezzi/",
			Validate: func(s string) error {
				if len(s) == 0 {
					return errors.New("cannot be empty")
				}
				return nil
			},
			Value: &module,
		},
		{
			Title:       "path",
			Placeholder: "somewhere/to/put/project",
			Prompt:      "~/tmp/",
			Focus:       true,
			Value:       &projectPath,
		},
		{
			Title:  "template",
			IsList: true,
			Value:  &template,
		},
	}

	err := CreateForm(fieldDefs)
	if err != nil {
		return nil, err
	}

	return cli.NewArguments(module, projectPath, template), nil
}

func Form(fields ...Field) error {
	m, err := newModel(fields...)
	if err != nil {
		return err
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
	if m.aborted {
		return errors.New("was aborted")
	}

	return nil
}
