package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/thejezzi/mkgo/internal/args"
)

func SummaryFromArguments(args *args.Arguments) string {
	projectDir := args.Path()
	makefile := args.CreateMakefile()
	tmpl := strings.ToLower(args.Template())

	var steps []string
	steps = append(steps, helpStep("cd "+projectDir, "Change to your project directory"))

	if makefile {
		steps = append(
			steps,
			helpStep("make build", "Build your project"),
			helpStep("make test", "Run all tests"),
		)
	} else {
		steps = append(
			steps,
			helpStep("go test ./...", "Run all tests"),
			helpStep("go run main.go", "Run your application"),
		)
	}

	if tmpl == "git" || args.InitGit() {
		steps = append(steps, helpStep("git commit -m 'initial'", "Commit your changes"))
	}

	var enumeratedSteps strings.Builder
	enumeratedSteps.WriteString("\n") // blank line above Next steps
	for i, s := range steps {
		enumeratedSteps.WriteString(fmt.Sprintf("  %d. %s", i+1, s))
	}

	successMsg := lipgloss.NewStyle().
		Foreground(lipgloss.Color("42")).
		Bold(true).
		Render("Project created successfully!")

	nextStepsTitle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("42")).
		Bold(true).
		Render("Next steps:")

	motivational := lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")).
		Render("Have fun building your project! 🚀")

	return strings.Join([]string{
		"",
		successMsg,
		nextStepsTitle,
		enumeratedSteps.String(),
		motivational,
	}, "\n")
}

func helpStep(cmd, desc string) string {
	return "  " + lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Render(cmd) + "  " + desc + "\n"
}
