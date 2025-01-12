package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

type inputModel struct {
	title            string
	titleStyle       lipgloss.Style
	description      string
	descriptionStyle lipgloss.Style
	inner            textinput.Model
	prompt           string
	promptIndex      int
	promptList       []string
	promptStyle      lipgloss.Style

	validation func(string) error
}

func NewInputModel() inputModel {
	return inputModel{
		titleStyle:       titleStyle,
		descriptionStyle: helpStyle,
		promptStyle:      focusedStyle,
		inner:            textinput.New(),
		promptList:       make([]string, 0),
	}
}
