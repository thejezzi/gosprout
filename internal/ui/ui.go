package ui

import (
	"errors"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/thejezzi/gosprout/cmd/sprout/cli"
)

func New() (*cli.Arguments, error) {
	var module, projectPath, template, gitRepo string

	fieldDefs := []FieldDef{
		{
			Title:       "Module",
			Description: "The name of your Go module",
			Placeholder: "github.com/your-username/your-project",
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
			Title:       "Path",
			Description: "The directory where your project will be created",
			Placeholder: "~/projects/my-go-app",
			Prompt:      "~/tmp/",
			Focus:       true,
			Value:       &projectPath,
		},
		{
			Title:       "Template",
			Description: "Choose a template to quickly set up your project structure",
			IsList:      true,
			Value:       &template,
		},
		{
			Title:       "Git Repository",
			Description: "Specify a Git repository to clone from (only for 'Git' template)",
			Placeholder: "github.com/user/repo",
			Prompt:      "https://",
			Value:       &gitRepo,
			Hide: func() bool {
				return template != "Git"
			},
		},
	}

	err := CreateForm(fieldDefs)
	if err != nil {
		return nil, err
	}

	return cli.NewArguments(module, projectPath, template, gitRepo), nil
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
