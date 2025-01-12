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

func Run() (module, projectPath, error) {
	m := newModel()
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
