package structure

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thejezzi/gosprout/internal/util"
)

func TestCreateGoMod(t *testing.T) {
	gomod, err := newGoMod("testmod")
	assert.NoError(t, err)
	assert.Greater(t, len(gomod), 0)

	gomodLines := strings.Split(string(gomod), string('\n'))
	assert.Equal(t, gomodLines[0], "module testmod")
	assert.Len(t, gomodLines[1], 0)
	assert.Contains(t, gomodLines[2], "go 1.")
}

func TestWriteGoModToFile(t *testing.T) {
	gomod, err := newGoMod("testmod")
	assert.NoError(t, err)

	tmpDir := filepath.Join(
		os.TempDir(),
		fmt.Sprintf("gosprout_test_%s", util.RandomString(6)),
	)

	err = gomod.WriteToFile(tmpDir)
	assert.NoError(t, err)
	fileInfo, err := os.Stat(filepath.Join(tmpDir, _gomodFileName))
	assert.NoError(t, err)
	assert.False(t, fileInfo.IsDir())
	_ = os.RemoveAll(tmpDir)
}

func TestGoVersion(t *testing.T) {
	ver := goVersion()
	assert.NotContains(t, ver, "go")
	assert.Contains(t, ver, "1.")
}

func TestCreateMainGo(t *testing.T) {
	tmpDir := filepath.Join(
		os.TempDir(),
		fmt.Sprintf("gosprout_test_%s", util.RandomString(6)),
	)
	err := newMainGo(tmpDir)
	assert.NoError(t, err)
	fileInfo, err := os.Stat(filepath.Join(tmpDir, _mainGoFileName))
	assert.NoError(t, err)
	assert.False(t, fileInfo.IsDir())
	_ = os.RemoveAll(tmpDir)
}
