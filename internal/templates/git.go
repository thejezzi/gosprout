package templates

import (
	"fmt"
	"github.com/thejezzi/gosprout/cmd/sprout/cli"
	"github.com/thejezzi/gosprout/internal/git"
	"github.com/thejezzi/gosprout/internal/model"
	"github.com/thejezzi/gosprout/internal/structure"
)

func gitCreate(args *cli.Arguments) error {
	if args.GitRepo == "" {
		return fmt.Errorf("git repository URL cannot be empty")
	}
	if err := git.Clone(args.GitRepo, args.Path()); err != nil {
		return err
	}
	if err := git.Reinit(args.Path()); err != nil {
		return err
	}
	return structure.ReplaceModuleName(args.Path(), args.Name())
}

var Git = model.NewTemplate(
	"Git",
	"Create a project from a git repository",
	gitCreate,
)
