package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func List() Field {
	return newListModel()
}

type listModel struct {
	listTitle string
	items     []list.Item
	focused   bool

	inner *list.Model
}

func newListModel() *listModel {
	newlist := list.New(
		[]list.Item{
			item{title: "Raspberry Pi’s", desc: "I have ’em all over my house"},
			item{title: "Nutella", desc: "It's good on toast"},
			item{title: "Bitter melon", desc: "It cools you down"},
			item{title: "Nice socks", desc: "And by that I mean socks without holes"},
			item{title: "Eight hours of sleep", desc: "I had this once"},
			item{title: "Cats", desc: "Usually"},
			item{title: "Plantasia, the album", desc: "My plants love it too"},
			item{title: "Pour over coffee", desc: "It takes forever to make though"},
			item{title: "VR", desc: "Virtual reality...what is there to say?"},
			item{title: "Noguchi Lamps", desc: "Such pleasing organic forms"},
			item{title: "Linux", desc: "Pretty much the best OS"},
			item{title: "Business school", desc: "Just kidding"},
			item{title: "Pottery", desc: "Wet clay is a great feeling"},
			item{title: "Shampoo", desc: "Nothing like clean hair"},
			item{title: "Table tennis", desc: "It’s surprisingly exhausting"},
			item{title: "Milk crates", desc: "Great for packing in your extra stuff"},
			item{title: "Afternoon tea", desc: "Especially the tea sandwich part"},
			item{title: "Stickers", desc: "The thicker the vinyl the better"},
			item{title: "20° Weather", desc: "Celsius, not Fahrenheit"},
			item{title: "Warm light", desc: "Like around 2700 Kelvin"},
			item{title: "The vernal equinox", desc: "The autumnal equinox is pretty good too"},
			item{title: "Gaffer’s tape", desc: "Basically sticky fabric"},
			item{title: "Terrycloth", desc: "In other words, towel fabric"},
		},
		list.NewDefaultDelegate(),
		200,
		10,
	)

	newlist.Title = ""

	nopaddingNewLine := lipgloss.NewStyle().Padding(0, 0, 1, 0)
	newlist.Styles.Title = noStyle
	newlist.Styles.TitleBar = noStyle
	newlist.Styles.StatusBar = nopaddingNewLine
	newlist.Styles.PaginationStyle = nopaddingNewLine
	newlist.Styles.HelpStyle = nopaddingNewLine
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

func (lm *listModel) Value(*string) Field {
	return lm
}

func (lm *listModel) Placeholder(s string) Field {
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
	return current.title
}
