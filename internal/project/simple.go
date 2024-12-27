package project

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"golang.org/x/mod/modfile"
)

type projectError struct {
	msg string
}

func (pe projectError) Error() string {
	return pe.msg
}

func perrorf(msg string, args ...any) *projectError {
	return &projectError{fmt.Sprintf(msg, args...)}
}

func NewSimple(module string) error {
	if module == "" {
		return errors.New("module name cannot be empty")
	}

	if err := os.Mkdir(module, 0755); err != nil {
		return perrorf("cannot create new project folder: %v", err)
	}

	if err := CreateGoMod(module); err != nil {
		return perrorf("module file: %v", err)
	}

	if err := CreateMainGo(module); err != nil {
		return perrorf("main file: %v", err)
	}

	return nil
}

// CreateGoMod creates a new go.mod file programmatically with the specified module name.
func CreateGoMod(moduleName string) error {
	if moduleName == "" {
		return fmt.Errorf("module name cannot be empty")
	}

	goVersion := runtime.Version()
	if len(goVersion) > 2 && goVersion[:2] == "go" {
		goVersion = goVersion[2:] // Strip "go" prefix
	}

	modFile := &modfile.File{}

	if err := modFile.AddModuleStmt(moduleName); err != nil {
		return fmt.Errorf("could not add module statement: %v", err)
	}

	if err := modFile.AddGoStmt(goVersion); err != nil {
		return fmt.Errorf("failed to add Go version directive: %w", err)
	}

	modData, err := modFile.Format()
	if err != nil {
		return fmt.Errorf("failed to format modfile: %w", err)
	}

	if err := os.WriteFile(fmt.Sprintf("%s/go.mod", moduleName), modData, 0644); err != nil {
		return fmt.Errorf("failed to write go.mod file: %w", err)
	}

	return nil
}

// CreateMainGo creates a new main.go file with a basic "Hello, World!" program.
func CreateMainGo(moduleName string) error {
	// Validate the module name
	if moduleName == "" {
		return perrorf("module name cannot be empty")
	}

	// Define the content of the main.go file
	mainGoContent := `package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
`

	// Convert module name to a directory-friendly path
	moduleDir := filepath.FromSlash(moduleName)

	// Define the path to the main.go file
	mainGoPath := filepath.Join(moduleDir, "main.go")

	// Write the main.go file
	if err := os.WriteFile(mainGoPath, []byte(mainGoContent), 0644); err != nil {
		return fmt.Errorf("failed to write main.go file: %w", err)
	}

	return nil
}
