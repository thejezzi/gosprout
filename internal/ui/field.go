package ui

import tea "github.com/charmbracelet/bubbletea"

type Field interface {
	Title(string) Field
	Description(string) Field
	Prompt(...string) Field
	FocusOnStart() Field
	Value(*string) Field
	Placeholder(s string) Field

	getTitle() string
	focus() tea.Cmd
	blur()
	update(tea.Msg) tea.Cmd
	isFocused() bool
	render() string
}
