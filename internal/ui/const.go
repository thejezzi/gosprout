package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	focusButtonStyle    = lipgloss.NewStyle().Background(lipgloss.Color("205")).Foreground(lipgloss.Color("0"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	listFocusedStyle   = lipgloss.NewStyle().Background(lipgloss.Color("#0066cc")).Foreground(lipgloss.Color("255"))
	listUnfocusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))

	focusedButton = focusButtonStyle.Render("  Submit  ")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)
