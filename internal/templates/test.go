package templates

import (
	"github.com/thejezzi/gosprout/cmd/sprout/cli"
	"github.com/thejezzi/gosprout/internal/model"
	"github.com/thejezzi/gosprout/internal/structure"
)

func testCreate(args *cli.Arguments) error {
	return structure.CreateNewModuleWithTest(args)
}

var Test = model.NewTemplate(
	"Test",
	"A cmd folder with a main_test.go file",
	testCreate,
)
