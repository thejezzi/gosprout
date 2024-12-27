package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/thejezzi/gosprout/internal"
)

type Field[T any] struct {
	name         string
	value        *T
	huhRepr      huh.Field
	descpription string
	validation   func(T) error
}

func NewField[T any](name string, value T) *Field[T] {
	f := Field[T]{name: name, value: &value}
	repr, ok := huhRepresentation(f.name, f.value, f.descpription)
	if ok {
		f.huhRepr = repr
	}
	return &f
}

func (f *Field[T]) Description(s string) *Field[T] {
	f.descpription = s
	return f
}

func huhRepresentation[T any](name string, value *T, descr string) (huh.Field, bool) {
	switch interface{}(*value).(type) {
	case string:
		return huh.NewInput().
			Title(name).
			Value(any(value).(*string)).Description(descr), true
	}

	return nil, false
}

func (f *Field[T]) String() string {
	if f.value == nil {
		return ""
	}
	return fmt.Sprintf("%v", *f.value)
}

type options struct {
	moduleName *Field[string]
	foldername *Field[string]
}

func buildPathPrefix() string {
	pathSlice := make([]string, 2)
	pathSlice[0] = internal.EnsureEnv("SPROUT_SCMPROVIDER", "github.com")
	pathSlice[1] = internal.EnsureEnv("SPROUT_USERNAME", "you")

	pathSlice = internal.DiscardEmptyElements(pathSlice)
	pathPrefix := strings.Join(pathSlice, "/")
	if len(pathPrefix) > 0 {
		pathPrefix += string(filepath.Separator)
	}
	return pathPrefix
}

func main() {
	opts := options{
		moduleName: NewField("module", buildPathPrefix()).Description("Wow"),
		foldername: NewField("folder", ""),
	}

	form := huh.NewForm(
		huh.NewGroup(opts.moduleName.huhRepr, opts.foldername.huhRepr),
	)

	_ = form.Run()
	fmt.Println(opts.moduleName.String(), opts.foldername.String())
}
