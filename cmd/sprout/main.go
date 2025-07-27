package main

import (
	"fmt"
	"os"
	"os/exec"

	argsPkg "github.com/thejezzi/gosprout/internal/args"
	"github.com/thejezzi/gosprout/internal/template"
	"github.com/thejezzi/gosprout/internal/ui"
	"github.com/thejezzi/gosprout/internal/util"
)

// newArgumentsFromUI converts a ui.UI to *argsPkg.Arguments
func newArgumentsFromUI(f ui.Form) *argsPkg.Arguments {
	// Converts a ui.Form to *argsPkg.Arguments
	if f == nil {
		return nil
	}
	return argsPkg.NewArguments(
		f.GetModule(),
		f.GetPath(),
		f.GetTemplate(),
		f.GetGitRepo(),
		f.GetCreateMakefile(),
		f.GetInitGit(),
	)
}

// getArguments returns Arguments from flags or UI
func getArguments() (*argsPkg.Arguments, error) {
	if len(os.Args) > 1 {
		return argsPkg.Flags()
	}
	form, err := ui.NewForm()
	if err != nil {
		return nil, err
	}
	return newArgumentsFromUI(form), nil
}

// printSummary prints a styled summary of the created project and next steps
func printSummary(args *argsPkg.Arguments) {
	title := ui.TitleStyle.Render("Project Created!")

	summary := ui.AppStyle.Render(
		ui.HelpStyle.Render("Module Name     : ") + args.Name() + "\n" +
			ui.HelpStyle.Render("Project Path    : ") + args.Path() + "\n" +
			ui.HelpStyle.Render("Template        : ") + args.Template() + "\n" +
			ui.HelpStyle.Render("Git Repository  : ") + args.GitRepo() + "\n" +
			ui.HelpStyle.Render("Makefile        : ") + fmt.Sprintf("%v", args.CreateMakefile()) + "\n",
	)

	// Template-specific next steps
	var nextSteps []string
	nextSteps = append(nextSteps, fmt.Sprintf("%s cd %s", ui.HelpStyle.Render("1."), args.Path()))

	switch args.Template() {
	case template.Git.Name:
		nextSteps = append(nextSteps, fmt.Sprintf("%s Check your .git configuration and remote origin.", ui.HelpStyle.Render("2.")))
		nextSteps = append(nextSteps, fmt.Sprintf("%s Start coding your project!", ui.HelpStyle.Render("3.")))
	case template.Test.Name:
		nextSteps = append(nextSteps, fmt.Sprintf("%s Run 'go test ./...' to verify your test setup.", ui.HelpStyle.Render("2.")))
		nextSteps = append(nextSteps, fmt.Sprintf("%s Start coding your project!", ui.HelpStyle.Render("3.")))
	case template.Simple.Name:
		nextSteps = append(nextSteps, fmt.Sprintf("%s Start coding your project!", ui.HelpStyle.Render("2.")))
	}
	if args.CreateMakefile() {
		nextSteps = append(nextSteps, fmt.Sprintf("%s Use 'make' to build or test your project.", ui.HelpStyle.Render("4.")))
	}

	fmt.Println("\n" + title)
	fmt.Println(summary)
	fmt.Println(ui.TitleStyle.Render("Next steps:"))
	for _, step := range nextSteps {
		fmt.Println(step)
	}
	fmt.Println("\n" + ui.HelpStyle.Render("Happy hacking! 🚀") + "\n")
}

func run() error {
	args, err := getArguments()
	if err != nil {
		return err
	}

	for _, t := range template.All {
		if t.Name == args.Template() {
			err := t.Create(args)
			if err == nil {
				if args.InitGit() {
					// Run git init in the project root
					cmd := exec.Command("git", "init")
					cmd.Dir = args.Path()
					_ = cmd.Run() // ignore error for now
				}
				printSummary(args)
			}
			return err
		}
	}

	return fmt.Errorf("template not found: %s", args.Template())
}

func main() {
	f, err := util.InitLogger()
	if err != nil {
		fmt.Printf("could not initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	if err := run(); err != nil {
		fmt.Printf("could not create new project: %v\n", err)
	}
}
