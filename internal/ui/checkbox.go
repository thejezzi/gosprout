package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type CheckboxField interface {
	Field
}

func Checkbox() CheckboxField {
	return newCheckboxModel()
}

type checkboxModel struct {
	title       string
	value       *bool
	focused     bool
	hidden      func() bool
	description string
}

func newCheckboxModel() *checkboxModel {
	return &checkboxModel{}
}

func (cm *checkboxModel) Title(s string) Field {
	cm.title = s
	return cm
}

func (cm *checkboxModel) Description(s string) Field {
	cm.description = s
	return cm
}

func (cm *checkboxModel) RotationDescription(string) Field {
	return cm
}

func (cm *checkboxModel) Prompt(...string) Field {
	return cm
}

func (cm *checkboxModel) FocusOnStart() Field {
	cm.focused = true
	return cm
}

func (cm *checkboxModel) Value(v *string) Field {
	// This is a hack to satisfy the Field interface.
	return cm
}

func (cm *checkboxModel) SetValue(v *bool) *checkboxModel {
	cm.value = v
	return cm
}

func (cm *checkboxModel) Placeholder(s string) Field {
	return cm
}

func (cm *checkboxModel) Validate(func(string) error) Field {
	return cm
}

func (cm *checkboxModel) DisablePromptRotation() Field {
	return cm
}

func (cm *checkboxModel) getTitle() string {
	return cm.title
}

func (cm *checkboxModel) focus() tea.Cmd {
	cm.focused = true
	return nil
}

func (cm *checkboxModel) blur() {
	cm.focused = false
}

func (cm *checkboxModel) update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case " ", "enter":
			*cm.value = !*cm.value
		}
	}
	return nil
}

func (cm *checkboxModel) isFocused() bool {
	return cm.focused
}

func (cm *checkboxModel) render() string {
	var s strings.Builder

	checkboxView := "☐ "
	if *cm.value {
		checkboxView = "✓ "
	}

	checkboxView += cm.description

	if cm.focused {
		s.WriteString(focusedStyle.Render(checkboxView))
	} else {
		s.WriteString(checkboxView)
	}

	s.WriteString("\n\n")

	return s.String()
}

func (cm *checkboxModel) SetHide(hide func() bool) {
	cm.hidden = hide
}

func (cm *checkboxModel) IsHidden() bool {
	if cm.hidden == nil {
		return false
	}
	return cm.hidden()
}
