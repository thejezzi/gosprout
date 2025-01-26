package ui

import (
	"errors"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type FieldTitle string

const (
	FieldTitleModule FieldTitle = "module"
	FieldTitlePath   FieldTitle = "path"
)

var ErrFieldIsNotAnInputModel = errors.New("field is not an input model")

type model struct {
	focusIndex int
	fields     []*inputModel
	cursorMode cursor.Mode
	aborted    bool
}

func newModel(fields ...Field) (*model, error) {
	m := model{
		fields: make([]*inputModel, len(fields)),
	}

	for i, field := range fields {
		input, ok := field.(*inputModel)
		if !ok {
			return nil, ErrFieldIsNotAnInputModel
		}
		if i == 0 {
			input.Focus()
			input.SetInnerTextStyle(focusedStyle)
		}

		input.SetInnerCursorStyle(cursorStyle)
		input.CharLimit(256)
		m.fields[i] = input
	}

	return &m, nil
}

var errFieldDoesNotExist = errors.New("field does not exist")

func (m *model) findFieldByTitle(t FieldTitle) (*inputModel, error) {
	for i := range m.fields {
		if m.fields[i].title == string(t) {
			return m.fields[i], nil
		}
	}
	return nil, errFieldDoesNotExist
}

func (m *model) Init() tea.Cmd {
	return textinput.Blink
}

func (m *model) setAllCursorsBlink() {
	for i := range m.fields {
		m.fields[i].SetInnerCursorMode(cursor.CursorBlink)
	}
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		cmd := m.updateInputs(msg)
		return m, cmd
	}

	switch keyMsg.String() {
	case "ctrl+c", "esc":
		m.aborted = true
		return m, tea.Quit

	case "ctrl+r":
		m.fields[m.focusIndex].RotatePrompt()
		cmd := m.updateInputs(msg)
		return m, cmd

	// Set focus to next input
	case "tab", "shift+tab", "enter", "up", "down":
		return m.focusNext(keyMsg)
	}

	cmd := m.updateInputs(msg)
	return m, cmd
}

func (m *model) focusNext(msg tea.KeyMsg) (*model, tea.Cmd) {
	s := msg.String()
	if s == "enter" && m.focusIndex == len(m.fields) {
		return m, tea.Quit
	}

	if s == "up" || s == "shift+tab" {
		m.focusIndex--
	} else {
		m.setAllCursorsBlink()
		m.focusIndex++
	}

	if m.focusIndex > len(m.fields) {
		m.focusIndex = 0
	} else if m.focusIndex < 0 {
		m.focusIndex = len(m.fields)
	}

	return m, tea.Batch(m.evaluateFocusStyles()...)
}

func (m *model) evaluateFocusStyles() []tea.Cmd {
	cmds := make([]tea.Cmd, len(m.fields))
	for i := 0; i <= len(m.fields)-1; i++ {
		if i == m.focusIndex {
			// Set focused state
			cmds[i] = m.fields[i].Focus()
			m.fields[i].SetInnerPromptStyle(focusedStyle)
			m.fields[i].SetInnerTextStyle(focusedStyle)
			continue
		}
		// Remove focused state
		m.fields[i].Blur()
		m.fields[i].SetInnerPromptStyle(noStyle)
		m.fields[i].SetInnerPromptStyle(noStyle)
	}
	return cmds
}

// updateInputs updates the inner textinput elements and nothing more.
func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.fields))
	for i, input := range m.fields {
		cmds[i] = input.UpdateInner(msg)
		m.fields[i] = input
	}
	return tea.Batch(cmds...)
}

func (m *model) View() string {
	var b strings.Builder
	for _, input := range m.fields {
		b.WriteString(input.render())
	}

	button := &blurredButton
	if m.focusIndex == len(m.fields) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "%s\n\n", *button)
	b.WriteString(helpStyle.Render("ctrl+r to change module prefix"))

	return b.String()
}
