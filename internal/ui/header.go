package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type HeaderField interface {
	Field
}

func Header() HeaderField {
	return &headerModel{}
}

type headerModel struct {
	title  string
	hidden func() bool
}

func (hm *headerModel) Title(s string) Field {
	hm.title = s
	return hm
}

func (hm *headerModel) Description(string) Field {
	return hm
}

func (hm *headerModel) RotationDescription(string) Field {
	return hm
}

func (hm *headerModel) Prompt(...string) Field {
	return hm
}

func (hm *headerModel) FocusOnStart() Field {
	return hm
}

func (hm *headerModel) Value(*string) Field {
	return hm
}

func (hm *headerModel) Placeholder(s string) Field {
	return hm
}

func (hm *headerModel) Validate(func(string) error) Field {
	return hm
}

func (hm *headerModel) DisablePromptRotation() Field {
	return hm
}

func (hm *headerModel) getTitle() string {
	return hm.title
}

func (hm *headerModel) focus() tea.Cmd {
	return nil
}

func (hm *headerModel) blur() {
}

func (hm *headerModel) update(msg tea.Msg) tea.Cmd {
	return nil
}

func (hm *headerModel) isFocused() bool {
	return false
}

func (hm *headerModel) render() string {
	return TitleStyle.Render(hm.title) + "\n"
}

func (hm *headerModel) SetHide(hide func() bool) {
	hm.hidden = hide
}

func (hm *headerModel) IsHidden() bool {
	if hm.hidden == nil {
		return false
	}
	return hm.hidden()
}
