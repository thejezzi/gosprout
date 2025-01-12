package ui

type Input interface {
	Title(string) Input
	Description(string) Input
	Value(value *any) Input
	FocusOnStart() Input
	WithPrefixes([]string) Input
}

type Form interface {
	Run() error
}

func NewForm(fields ...Input) Form {
}
