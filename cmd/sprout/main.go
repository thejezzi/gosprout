package main

import (
	"errors"
	"flag"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/thejezzi/gosprout/internal/project"
	"github.com/thejezzi/gosprout/internal/util"
)

func buildPathPrefix() string {
	pathSlice := make([]string, 2)
	pathSlice[0] = util.EnsureEnv("SPROUT_SCMPROVIDER", "github.com")
	pathSlice[1] = util.EnsureEnv("SPROUT_USERNAME", "you")

	pathSlice = util.DiscardEmptyElements(pathSlice)
	pathPrefix := strings.Join(pathSlice, "/")
	if len(pathPrefix) > 0 {
		pathPrefix += string(filepath.Separator)
	}
	return pathPrefix
}

func validateModuleName(s string) error {
	if len(strings.Split(s, " ")) > 1 {
		return errors.New("must be one word")
	}

	if strings.Index(s, "/") > 0 {
		moduleSlice := strings.Split(s, "/")
		if len(moduleSlice) <= 2 {
			return errors.New("if you provide a module path min length is 3")
		}

		if moduleSlice[2] == "" {
			return errors.New("you have to provide a module name")
		}
	}

	return nil
}

func runUI(module *string) {
	*module = buildPathPrefix()
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("module").
				Value(module).
				Validate(validateModuleName),
		),
	)

	_ = form.Run()
}

func main() {
	module := flag.String("module", "", "provide a module title")
	flag.Parse()

	if *module == "" {
		runUI(module)
	}

	if err := project.NewSimple(*module); err != nil {
		fmt.Println(err)
	}
}
