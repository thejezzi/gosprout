package git

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Reinit(path string) error {
	if err := os.RemoveAll(filepath.Join(path, ".git")); err != nil {
		return fmt.Errorf("could not remove .git directory: %w", err)
	}

	cmd := exec.Command("git", "init")
	cmd.Dir = path
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("could not re-initialize git repository: %w", err)
	}
	return nil
}

func Clone(repoURL, path string) error {
	cmd := exec.Command("git", "clone", repoURL, path)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("could not clone repository: %w", err)
	}
	return nil
}
