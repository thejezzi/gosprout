package ui

import (
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type InputField interface {
	Title(string) InputField
	Description(string) InputField
	Prompt(...string) InputField
	FocusOnStart() InputField
	Value(*string) InputField
	Placeholder(s string) InputField
}

func Input() InputField {
	im := newInputModel()
	return &im
}

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
	focusOnStart     bool

	value *string

	validation func(string) error
}

func newInputModel() inputModel {
	im := inputModel{
		titleStyle:       titleStyle,
		descriptionStyle: helpStyle,
		promptStyle:      focusedStyle,
		inner:            textinput.New(),
		promptList:       make([]string, 0),
	}

	im.AppendPrompts("")
	im.inner.Prompt = ""
	im.inner.Cursor.SetMode(cursor.CursorBlink)
	return im
}

func (im *inputModel) Title(s string) InputField {
	im.title = s
	return im
}

func (im *inputModel) Description(desc string) InputField {
	im.description = desc
	return im
}

func (im *inputModel) FocusOnStart() InputField {
	im.Focus()
	im.SetInnerCursorMode(cursor.CursorHide)
	return im
}

func (im *inputModel) Prompt(prompts ...string) InputField {
	im.promptList = append(im.promptList, prompts...)
	return im
}

func (im *inputModel) Value(v *string) InputField {
	im.value = v
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

func (im *inputModel) Placeholder(p string) InputField {
	im.inner.Placeholder = p
	return im
}

func (im *inputModel) RotatePrompt() {
	if len(im.promptList) == 0 {
		return
	}
	im.promptIndex++

	if im.promptIndex > len(im.promptList)-1 {
		im.promptIndex = 0
	}
	im.prompt = im.promptList[im.promptIndex]
}

func (im *inputModel) AppendPrompts(prompts ...string) {
	im.promptList = append(im.promptList, prompts...)
	im.prompt = im.promptList[im.promptIndex]
}

func (im *inputModel) SetInnerCursorStyle(s lipgloss.Style) {
	im.inner.Cursor.Style = s
}

func (im *inputModel) UpdateInner(msg tea.Msg) tea.Cmd {
	updated, cmd := im.inner.Update(msg)
	*im.value = im.prompt + updated.Value()
	im.inner = updated
	return cmd
}

func (im *inputModel) SetInnerPromptStyle(s lipgloss.Style) {
	im.inner.PromptStyle = s
}

func (im *inputModel) Blur() {
	im.inner.Blur()
}
