package ui

// Hidable defines the contract for UI elements that can be hidden.
type Hidable interface {
	// SetHide sets the function that determines the field's visibility.
	SetHide(func() bool)
	// IsHidden returns true if the field should be hidden.
	IsHidden() bool
}
