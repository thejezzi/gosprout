package ui

import (
	"errors"
	"fmt"
	"os"
	"path"

	tea "github.com/charmbracelet/bubbletea"
)

type (
	module      = string
	projectPath = string
)

type UiOpt func(m *model) error

func WithPrefixes(fieldTitle FieldTitle, prefixes ...string) UiOpt {
	return func(m *model) error {
		field, err := m.findFieldByTitle(fieldTitle)
		if err != nil {
			return fmt.Errorf("add prefixes to %s: %w", string(fieldTitle), err)
		}
		field.AppendPrompts(prefixes...)
		return nil
	}
}

func Run(opts ...UiOpt) (module, projectPath, error) {
	m := newModel()
	for _, opt := range opts {
		if err := opt(m); err != nil {
			return "", "", err
		}
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
	if m.aborted {
		return "", "", errors.New("was aborted")
	}

	moduleInput := m.inputs[0]
	moduleValue := path.Clean(path.Join(moduleInput.prompt, moduleInput.inner.Value()))
	projectPathValue := m.inputs[1].inner.Value()

	return moduleValue, projectPathValue, nil
}
