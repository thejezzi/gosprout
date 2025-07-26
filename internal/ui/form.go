package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/thejezzi/gosprout/internal/template"
)

type FieldDef struct {
	Title       string
	Description string
	Placeholder string
	Prompt      string
	Focus       bool
	Validate    func(string) error
	Value       *string
	IsList      bool
	Hide        func() bool
}

func createInputField(fd FieldDef) Field {
	input := Input().
		Title(fd.Title).
		Description(fd.Description).
		Placeholder(fd.Placeholder).
		Prompt(fd.Prompt).
		Validate(fd.Validate).
		Value(fd.Value)
	input.SetHide(fd.Hide)
	if fd.Focus {
		input.FocusOnStart()
	}
	return input
}

func createListField(fd FieldDef) Field {
	items := make([]list.Item, len(template.All))
	for i, t := range template.All {
		items[i] = t
	}
	list := List().SetItems(items...)
	return list.Title(fd.Title).Value(fd.Value)
}

func CreateForm(fieldDefs []FieldDef) error {
	fields := make([]Field, len(fieldDefs))
	for i, fd := range fieldDefs {
		if fd.IsList {
			fields[i] = createListField(fd)
			continue
		}
		fields[i] = createInputField(fd)
	}
	return Form(fields...)
}
