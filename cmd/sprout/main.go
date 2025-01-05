package main

import (
	"flag"
	"log"

	"github.com/thejezzi/gosprout/internal"
)

func main() {
	moduleName := flag.String(
		"module",
		"testproj",
		"your module name or path like github.com/you/proj",
	)

	modulePath := flag.String(
		"path",
		"testproj",
		"the path to put all the files",
	)

	template := flag.String(
		"template",
		"simple",
		"specify a template to avoid some boilerplate setup",
	)

	flag.Parse()

	if err := internal.CreateNewModule(*modulePath, *moduleName, *template); err != nil {
		log.Fatalf("failed: %v", err)
	}
}
