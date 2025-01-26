package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func Input() Field {
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

func (im *inputModel) Title(s string) Field {
	im.title = s
	return im
}

func (im *inputModel) Description(desc string) Field {
	im.description = desc
	return im
}

func (im *inputModel) FocusOnStart() Field {
	im.Focus()
	im.SetInnerCursorMode(cursor.CursorHide)
	return im
}

func (im *inputModel) Prompt(prompts ...string) Field {
	im.promptList = append(im.promptList, prompts...)
	return im
}

func (im *inputModel) Value(v *string) Field {
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

func (im *inputModel) Placeholder(p string) Field {
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

func (im *inputModel) render() string {
	b := strings.Builder{}
	// title
	b.WriteString(im.titleStyle.Render(im.title))
	b.WriteRune('\n')
	// the actual input
	im.inner.TextStyle = noStyle
	b.WriteString("> ")
	b.WriteString(im.promptStyle.Render(im.prompt))
	b.WriteString(im.inner.View())
	b.WriteRune('\n')
	if len(im.description) > 0 {
		b.WriteString(im.descriptionStyle.Render(im.description))
		b.WriteRune('\n')
	}
	b.WriteRune('\n')

	return b.String()
}
