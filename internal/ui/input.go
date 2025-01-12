package ui

import (
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
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
	im := inputModel{
		titleStyle:       titleStyle,
		descriptionStyle: helpStyle,
		promptStyle:      focusedStyle,
		inner:            textinput.New(),
		promptList:       make([]string, 0),
	}

	im.inner.Prompt = ""
	im.inner.Cursor.SetMode(cursor.CursorBlink)
	return im
}

func (im *inputModel) CharLimit(n int) {
	im.inner.CharLimit = n
}

func (im *inputModel) SetInnerCursorMode(mode cursor.Mode) tea.Cmd {
	return im.inner.Cursor.SetMode(mode)
}

func (im *inputModel) SetInnerTextStyle(s lipgloss.Style) {
	im.inner.TextStyle = s
}

func (im *inputModel) Focus() tea.Cmd {
	return im.inner.Focus()
}

func (im *inputModel) SetPlaceholder(p string) {
	im.inner.Placeholder = p
}

func (im *inputModel) RotatePrompt() {
	im.promptIndex++

	if im.promptIndex > len(im.promptList)-1 {
		im.promptIndex = 0
	}
	im.prompt = im.promptList[im.promptIndex]
}

func (im *inputModel) Prompts(prompts ...string) {
	im.promptList = append(im.promptList, prompts...)
	im.prompt = im.promptList[im.promptIndex]
}

func (im *inputModel) SetInnerCursorStyle(s lipgloss.Style) {
	im.inner.Cursor.Style = s
}

func (im *inputModel) UpdateInner(msg tea.Msg) tea.Cmd {
	updated, cmd := im.inner.Update(msg)
	im.inner = updated
	return cmd
}

func (im *inputModel) SetInnerPromptStyle(s lipgloss.Style) {
	im.inner.PromptStyle = s
}

func (im *inputModel) Blur() {
	im.inner.Blur()
}
