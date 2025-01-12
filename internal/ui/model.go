package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	focusIndex int
	inputs     []inputModel
	cursorMode cursor.Mode
	aborted    bool
}

func initialModel() *model {
	m := model{
		inputs: make([]inputModel, 2),
	}

	var t inputModel
	for i := range m.inputs {
		t = NewInputModel()
		t.inner.Cursor.Style = cursorStyle
		t.inner.CharLimit = 256

		switch i {
		case 0:
			t.title = "Module"
			t.description = "Your module path that is used in the go mod file"
			t.inner.Placeholder = "module"
			t.inner.Focus()
			t.inner.TextStyle = focusedStyle
			t.inner.Prompt = ""
			t.promptList = append(t.promptList, "", "github.com/you/", "bitbucket.org/you/")
			t.prompt = t.promptList[t.promptIndex]
		case 1:
			t.title = "Path"
			t.description = "The path where to put your project"
			t.inner.Placeholder = "path"
			t.inner.Focus()
			t.inner.Cursor.SetMode(cursor.CursorHide)
			t.inner.CharLimit = 256
			t.inner.Prompt = ""
		}

		m.inputs[i] = t
	}

	return &m
}

func (m *model) Init() tea.Cmd {
	return textinput.Blink
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.aborted = true
			return m, tea.Quit

		case "ctrl+r":
			cFocused := m.inputs[m.focusIndex]
			cFocused.promptIndex++

			if cFocused.promptIndex > len(cFocused.promptList)-1 {
				cFocused.promptIndex = 0
			}
			cFocused.prompt = cFocused.promptList[cFocused.promptIndex]
			m.inputs[m.focusIndex] = cFocused

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()
			if s == "enter" && m.focusIndex == len(m.inputs) {
				return m, tea.Quit
			}

			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.setCursorModeBlink()
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].inner.Focus()
					m.inputs[i].inner.PromptStyle = focusedStyle
					m.inputs[i].inner.TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].inner.Blur()
				m.inputs[i].inner.PromptStyle = noStyle
				m.inputs[i].inner.TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *model) setCursorModeBlink() {
	for i := range m.inputs {
		m.inputs[i].inner.Cursor.SetMode(cursor.CursorBlink)
	}
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		innerInput, cmd := m.inputs[i].inner.Update(msg)
		cmds[i] = cmd
		m.inputs[i].inner = innerInput
	}

	return tea.Batch(cmds...)
}

func (m *model) View() string {
	var b strings.Builder

	for i := range m.inputs {
		currentInput := m.inputs[i]
		b.WriteString(renderInput(currentInput))
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "%s\n\n", *button)

	b.WriteString(helpStyle.Render("ctrl+r to change prompt"))

	return b.String()
}

func renderInput(input inputModel) string {
	b := strings.Builder{}
	// title
	b.WriteString(input.titleStyle.Render(input.title))
	b.WriteRune('\n')
	// the actual input
	input.inner.TextStyle = noStyle
	b.WriteString("> ")
	b.WriteString(input.promptStyle.Render(input.prompt))
	b.WriteString(input.inner.View())
	b.WriteRune('\n')
	if len(input.description) > 0 {
		b.WriteString(input.descriptionStyle.Render(input.description))
		b.WriteRune('\n')
	}
	b.WriteRune('\n')

	return b.String()
}
