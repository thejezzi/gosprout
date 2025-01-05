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

	flag.Parse()

	if err := internal.Simple(*modulePath, *moduleName); err != nil {
		log.Fatalf("failed: %v", err)
	}
}
