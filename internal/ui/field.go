package ui

type Field interface {
	Title(string) Field
	Description(string) Field
	Prompt(...string) Field
	FocusOnStart() Field
	Value(*string) Field
	Placeholder(s string) Field

	render() string
}
