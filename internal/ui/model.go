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
	fields     []Field
	cursorMode cursor.Mode
	aborted    bool
}

func newModel(fields ...Field) (*model, error) {
	m := model{
		fields: make([]Field, len(fields)),
	}

	for i, field := range fields {
		if i == 0 {
			field.focus()
		}
		m.fields[i] = field
	}

	return &m, nil
}

func (m *model) Init() tea.Cmd {
	return textinput.Blink
}

func (m *model) setAllCursorsBlink() {
	for _, field := range m.fields {
		input, ok := field.(*inputModel)
		if !ok {
			continue
		}
		input.SetInnerCursorMode(cursor.CursorBlink)
	}
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); !ok {
		return m.handleButKeyMsg(msg)
	}

	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		cmd := m.updateFields(msg)
		return m, cmd
	}

	switch keyMsg.String() {
	case "ctrl+c", "esc":
		m.aborted = true
		return m, tea.Quit

	case "ctrl+r":
		input, ok := m.fields[m.focusIndex].(*inputModel)
		if !ok {
			break
		}
		input.rotatePrompt()
		cmd := m.updateFields(msg)
		return m, cmd

	// Set focus to next input
	case "tab", "shift+tab", "up", "down":
		return m.focusNext(keyMsg)
	case "enter":
		if m.focusIndex == len(m.fields) {
			return m, tea.Quit
		}
	}

	cmd := m.updateFields(msg)
	return m, cmd
}

func (m *model) handleButKeyMsg(msg tea.Msg) (*model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		for _, field := range m.fields {
			if list, ok := field.(*listModel); ok {
				list.setWidth(msg.Width)
			}
		}
	}

	return m, m.updateFields(msg)
}

func (m *model) focusNext(msg tea.KeyMsg) (*model, tea.Cmd) {
	s := msg.String()

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
	for i, field := range m.fields {
		if i == m.focusIndex {
			cmds[i] = field.focus()
			continue
		}
		// Remove focused state
		field.blur()
	}
	return cmds
}

// updateFields updates the inner textinput elements and nothing more.
func (m *model) updateFields(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.fields))
	for i, field := range m.fields {
		if field.isFocused() {
			cmds[i] = field.update(msg)
		}
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
