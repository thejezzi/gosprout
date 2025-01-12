package structure

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"golang.org/x/mod/modfile"
)

const (
	_gomodFileName  = "go.mod"
	_mainGoFileName = "main.go"
	_mainGoContent  = `package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
`
)

type template = string

const (
	_templateSimple template = "simple"
)

type gomodData []byte

func newGoMod(moduleName string) (gomodData, error) {
	modFile := &modfile.File{}

	if err := modFile.AddModuleStmt(moduleName); err != nil {
		return nil, fmt.Errorf("could not add module statement: %v", err)
	}

	if err := modFile.AddGoStmt(goVersion()); err != nil {
		return nil, fmt.Errorf("failed to add Go version directive: %w", err)
	}

	modData, err := modFile.Format()
	if err != nil {
		return nil, fmt.Errorf("failed to format modfile: %w", err)
	}

	return modData, nil
}

func (gmd gomodData) WriteToFile(path string) error {
	cleanedPath := filepath.Clean(path)
	splitted := strings.Split(cleanedPath, "/")
	if splitted[len(splitted)-1] != _gomodFileName {
		splitted = append(splitted, _gomodFileName)
	}
	gomodPath := strings.Join(splitted, "/")

	dirPath := filepath.Dir(gomodPath)
	if err := ensureDir(dirPath); err != nil {
		return fmt.Errorf("could not create main.go file: %v", err)
	}
	if err := os.WriteFile(gomodPath, gmd, 0o644); err != nil {
		return fmt.Errorf("failed to write go.mod file: %w", err)
	}
	return nil
}

// type Options struct {
// 	path     string
// 	name     string
// 	template string
// }
//
// func NewOptions(moduleName, modulePath string, temp template) *Options {
// 	return &Options{
// 		path:     modulePath,
// 		name:     moduleName,
// 		template: temp,
// 	}
// }
//
// func (opts Options) Name() string {
// 	return opts.name
// }
//
// func (opts Options) Path() string {
// 	return opts.path
// }
//
// func (opts Options) Template() string {
// 	return opts.template
// }

type options interface {
	Name() string
	Path() string
	Template() string
}

func CreateNewModule(opts options) error {
	switch opts.Template() {
	case _templateSimple:
		return simple(opts)
	}
	return nil
}

func simple(opts options) error {
	if err := ensureDir(opts.Path()); err != nil {
		return err
	}

	gomod, err := newGoMod(opts.Name())
	if err != nil {
		return err
	}
	if err := gomod.WriteToFile(opts.Path()); err != nil {
		return err
	}

	basename := path.Base(opts.Name())
	cmdPath := filepath.Join(opts.Path(), "cmd", basename)
	if err := newMainGo(cmdPath); err != nil {
		return err
	}

	return nil
}

func ensureDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0o755)
	}
	return nil
}

func goVersion() string {
	version := runtime.Version()
	if len(version) > 2 && version[:2] == "go" {
		version = version[2:] // Strip "go" prefix
	}
	return version
}

func newMainGo(path string) error {
	cleanedPath := filepath.Clean(path)
	splitted := strings.Split(cleanedPath, "/")
	if splitted[len(splitted)-1] != _mainGoFileName {
		splitted = append(splitted, _mainGoFileName)
	}
	mainGoPath := strings.Join(splitted, "/")

	dirPath := filepath.Dir(mainGoPath)
	if err := ensureDir(dirPath); err != nil {
		return fmt.Errorf("could not create main.go file: %v", err)
	}

	if err := os.WriteFile(mainGoPath, []byte(_mainGoContent), 0o644); err != nil {
		return fmt.Errorf("failed to write main.go file: %w", err)
	}

	return nil
}
