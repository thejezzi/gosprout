package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"testing"
)

type dummyField struct{}

// Hidable
func (d *dummyField) SetHide(f func() bool) {}
func (d *dummyField) IsHidden() bool        { return false }

// Field interface
func (d *dummyField) Title(s string) Field               { return d }
func (d *dummyField) Description(s string) Field         { return d }
func (d *dummyField) RotationDescription(s string) Field { return d }
func (d *dummyField) Prompt(...string) Field             { return d }
func (d *dummyField) FocusOnStart() Field                { return d }
func (d *dummyField) Value(*string) Field                { return d }
func (d *dummyField) Placeholder(s string) Field         { return d }
func (d *dummyField) Validate(func(string) error) Field  { return d }
func (d *dummyField) DisablePromptRotation() Field       { return d }
func (d *dummyField) getTitle() string                   { return "dummy" }
func (d *dummyField) focus() tea.Cmd                     { return nil }
func (d *dummyField) blur()                              {}
func (d *dummyField) update(msg tea.Msg) tea.Cmd         { return nil }
func (d *dummyField) isFocused() bool                    { return true }
func (d *dummyField) render() string                     { return "Hello\nWorld\n" }

func TestModelLineCount(t *testing.T) {
	m, err := newModel(&dummyField{})
	if err != nil {
		t.Fatalf("failed to create model: %v", err)
	}
	_ = m.View()
	if m.LineCount() < 2 {
		t.Errorf("expected at least 2 lines, got %d", m.LineCount())
	}
}
