package ui

import "github.com/charmbracelet/bubbles/list"

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func List() Field {
	return &listModel{}
}

type listModel struct {
	listTitle string
	items     []list.Item

	inner list.Model
}

func newListModel() *listModel {
	return &listModel{
		listTitle: "MyList",
		inner: list.New(
			[]list.Item{
				item{title: "wow", desc: "wow"},
			},
			list.NewDefaultDelegate(),
			0,
			0,
		),
	}
}

func (lm *listModel) Title(string) Field {
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

func (lm *listModel) render() string {
	return "Wow"
}
