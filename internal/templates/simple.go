package templates

import (
	"github.com/thejezzi/gosprout/cmd/sprout/cli"
	"github.com/thejezzi/gosprout/internal/model"
	"github.com/thejezzi/gosprout/internal/structure"
)

func simpleCreate(args *cli.Arguments) error {
	return structure.CreateNewModule(args)
}

var Simple = model.NewTemplate(
	"Simple",
	"A simple structure with a cmd folder",
	simpleCreate,
)
