package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ListField interface {
	Field

	SetItems(...list.Item) ListField
}

type item interface {
	Title() string
	Description() string
	FilterValue() string
}

func List() ListField {
	return newListModel()
}

type listModel struct {
	listTitle string
	items     []list.Item
	focused   bool

	inner      *list.Model
	outerValue *string
}

func newListModel() *listModel {
	newlist := list.New(
		[]list.Item{},
		list.NewDefaultDelegate(),
		200,
		9,
	)

	// Remove List Title for seperate rendering of list title
	newlist.Title = ""
	nopaddingNewLine := lipgloss.NewStyle().Padding(0, 0, 1, 0)
	newlist.Styles.Title = noStyle
	newlist.Styles.TitleBar = noStyle
	newlist.Styles.StatusBar = nopaddingNewLine
	newlist.Styles.PaginationStyle = nopaddingNewLine
	newlist.Styles.HelpStyle = nopaddingNewLine

	// Remap clear filter keymap to ctrl+r
	newlist.KeyMap.ClearFilter.SetKeys("ctrl+r")
	newlist.KeyMap.ClearFilter.SetHelp("ctrl+r", "clear filter")

	// remove the quit keymap to prevent the user from unintentionally quitting
	// the programm and creating a project
	newlist.KeyMap.Quit.Unbind()

	return &listModel{
		listTitle: "MyList",
		inner:     &newlist,
	}
}

func (lm *listModel) Title(s string) Field {
	lm.listTitle = s
	return lm
}

func (lm *listModel) Description(string) Field {
	return lm
}

func (lm *listModel) Prompt(...string) Field {
	return lm
}

func (lm *listModel) FocusOnStart() Field {
	return lm
}

func (lm *listModel) Value(outer *string) Field {
	lm.outerValue = outer
	return lm
}

func (lm *listModel) Placeholder(s string) Field {
	return lm
}

func (lm *listModel) SetItems(items ...list.Item) ListField {
	lm.inner.SetItems(items)
	return lm
}

func (lm *listModel) Validate(func(string) error) Field {
	return lm
}

var docStyle = lipgloss.NewStyle().Margin(1, 2)

func (lm *listModel) render() string {
	if lm.inner == nil {
		return ""
	}
	v := strings.Builder{}

	if !lm.focused {
		v.WriteString(titleStyle.Render(lm.listTitle) + "\n")
		v.WriteString(lm.renderCurrentSelection())
		v.WriteRune('\n')
		v.WriteRune('\n')
		return v.String()
	}

	divider := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")). // Set color
		Width(50).                         // Set width
		Render(strings.Repeat("-", lm.inner.Width()) + "\n")
	v.WriteString(divider)
	v.WriteString("\n")
	v.WriteString(lm.renderTitle())
	v.WriteString("\n")
	v.WriteString(lm.inner.View())
	v.WriteString("\n")
	v.WriteString(divider)
	v.WriteString("\n")
	return v.String()
}

func (lm *listModel) renderTitle() string {
	title := "  " + lm.listTitle + "  "
	return focusButtonStyle.Render(title)
}

func (lm *listModel) renderCurrentSelection() string {
	return "> " + lm.value()
}

func (lm *listModel) blur() {
	lm.inner.Styles.Title = listUnfocusedStyle
	lm.focused = false
}

func (lm *listModel) focus() tea.Cmd {
	lm.inner.Styles.Title = listFocusedStyle
	lm.focused = true
	return nil
}

func (lm *listModel) getTitle() string {
	return lm.listTitle
}

func (lm *listModel) update(msg tea.Msg) tea.Cmd {
	if lm.inner == nil {
		return nil
	}
	updated, cmd := lm.inner.Update(msg)
	lm.inner = &updated
	*lm.outerValue = lm.value()
	return cmd
}

func (lm *listModel) setWidth(width int) {
	lm.inner.SetSize(width, lm.inner.Height())
}

func (lm *listModel) isFocused() bool {
	return lm.focused
}

func (lm *listModel) value() string {
	current, ok := lm.inner.SelectedItem().(item)
	if !ok {
		return ""
	}
	return current.Title()
}
