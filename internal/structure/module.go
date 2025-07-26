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
	_gomodFileName    = "go.mod"
	_mainGoFileName   = "main.go"
	_mainTestFileName = "main_test.go"
	_mainGoContent    = `package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
`
	_mainTestGoContent = `package main

import "testing"

func TestMain(t *testing.T) {
	t.Log("Hello, World!")
}
`
)

type template = string

const (
	_templateSimple = "simple"
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

type options interface {
	Name() string
	Path() string
	Template() string
}

func CreateNewModule(opts options) error {
	if err := simple(opts); err != nil {
		return err
	}
	return nil
}

func CreateNewModuleWithTest(opts options) error {
	if err := simple(opts); err != nil {
		return err
	}

	basename := path.Base(opts.Name())
	cmdPath := filepath.Join(opts.Path(), "cmd", basename)
	if err := newMainTestGo(cmdPath); err != nil {
		return err
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

func newMainTestGo(path string) error {
	cleanedPath := filepath.Clean(path)
	splitted := strings.Split(cleanedPath, "/")
	if splitted[len(splitted)-1] != _mainTestFileName {
		splitted = append(splitted, _mainTestFileName)
	}
	mainTestGoPath := strings.Join(splitted, "/")

	dirPath := filepath.Dir(mainTestGoPath)
	if err := ensureDir(dirPath); err != nil {
		return fmt.Errorf("could not create main_test.go file: %v", err)
	}

	if err := os.WriteFile(mainTestGoPath, []byte(_mainTestGoContent), 0o644); err != nil {
		return fmt.Errorf("failed to write main_test.go file: %w", err)
	}

	return nil
}

func ReplaceModuleName(path, newName string) error {
	goModPath := filepath.Join(path, _gomodFileName)
	goModBytes, err := os.ReadFile(goModPath)
	if err != nil {
		return fmt.Errorf("could not read go.mod file: %w", err)
	}

	modFile, err := modfile.Parse(goModPath, goModBytes, nil)
	if err != nil {
		return fmt.Errorf("could not parse go.mod file: %w", err)
	}

	if err := modFile.AddModuleStmt(newName); err != nil {
		return fmt.Errorf("could not add module statement: %v", err)
	}

	modData, err := modFile.Format()
	if err != nil {
		return fmt.Errorf("failed to format modfile: %w", err)
	}

	if err := os.WriteFile(goModPath, modData, 0o644); err != nil {
		return fmt.Errorf("failed to write go.mod file: %w", err)
	}

	return nil
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
