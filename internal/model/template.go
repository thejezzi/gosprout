package model

import "github.com/thejezzi/gosprout/cmd/sprout/cli"

func NewTemplate(name, description string, create func(args *cli.Arguments) error) Template {
	return Template{
		Name:        name,
		description: description,
		Create:      create,
	}
}

type Template struct {
	Name        string
	description string
	Create      func(args *cli.Arguments) error
}

func (t Template) Title() string {
	return t.Name
}

func (t Template) Description() string {
	return t.description
}

func (t Template) FilterValue() string {
	return t.Name
}
