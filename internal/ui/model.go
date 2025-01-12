package ui

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/thejezzi/gosprout/internal/util"
)

type model struct {
	focusIndex int
	inputs     []inputModel
	cursorMode cursor.Mode
	aborted    bool
}

func newModel() *model {
	m := model{
		inputs: make([]inputModel, 2),
	}

	projectPlaceholder, err := util.RandomSample(util.RandomProject)
	if err != nil {
		projectPlaceholder = "yourproject"
	}
	randomPath, err := util.RandomSample(util.RandomPath)
	if err != nil {
		randomPath = "/path/to/your/project"
	}
	pathPlaceholder := filepath.Join(randomPath, projectPlaceholder)

	var t inputModel
	for i := range m.inputs {
		t = NewInputModel()
		t.SetInnerCursorStyle(cursorStyle)
		t.CharLimit(256)

		switch i {
		case 0:
			t.title = "Module"
			t.description = "Your module path that is used in the go mod file"
			t.SetPlaceholder(projectPlaceholder)
			t.Focus()
			t.SetInnerTextStyle(focusedStyle)
			t.Prompts("", "github.com/you/", "bitbucket.org/you/")
		case 1:
			t.title = "Path"
			t.description = "The path where to put your project"
			t.SetPlaceholder(pathPlaceholder)
			t.Focus()
			t.SetInnerCursorMode(cursor.CursorHide)
		}

		m.inputs[i] = t
	}

	return &m
}

func (m *model) Init() tea.Cmd {
	return textinput.Blink
}

func (m *model) setAllCursorsBlink() {
	for i := range m.inputs {
		m.inputs[i].SetInnerCursorMode(cursor.CursorBlink)
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
		m.inputs[m.focusIndex].RotatePrompt()
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
	if s == "enter" && m.focusIndex == len(m.inputs) {
		return m, tea.Quit
	}

	if s == "up" || s == "shift+tab" {
		m.focusIndex--
	} else {
		m.setAllCursorsBlink()
		m.focusIndex++
	}

	if m.focusIndex > len(m.inputs) {
		m.focusIndex = 0
	} else if m.focusIndex < 0 {
		m.focusIndex = len(m.inputs)
	}

	return m, tea.Batch(m.evaluateFocusStyles()...)
}

func (m *model) evaluateFocusStyles() []tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))
	for i := 0; i <= len(m.inputs)-1; i++ {
		if i == m.focusIndex {
			// Set focused state
			cmds[i] = m.inputs[i].Focus()
			m.inputs[i].SetInnerPromptStyle(focusedStyle)
			m.inputs[i].SetInnerTextStyle(focusedStyle)
			continue
		}
		// Remove focused state
		m.inputs[i].Blur()
		m.inputs[i].SetInnerPromptStyle(noStyle)
		m.inputs[i].SetInnerPromptStyle(noStyle)
	}
	return cmds
}

// updateInputs updates the inner textinput elements and nothing more.
func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))
	for i, input := range m.inputs {
		cmds[i] = input.UpdateInner(msg)
		m.inputs[i] = input
	}
	return tea.Batch(cmds...)
}

func (m *model) View() string {
	var b strings.Builder
	for _, input := range m.inputs {
		b.WriteString(renderInput(input))
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
