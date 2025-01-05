package internal

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func randomString(length int) string {
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		randomType := rand.Intn(3)
		switch randomType {
		case 0:
			result[i] = byte(48 + rand.Intn(10))
		case 1:
			result[i] = byte(65 + rand.Intn(26))
		default:
			result[i] = byte(97 + rand.Intn(26))
		}
	}

	return string(result)
}

func TestCreateGoMod(t *testing.T) {
	gomod, err := createGoMod("testmod")
	assert.NoError(t, err)
	assert.Greater(t, len(gomod), 0)

	gomodLines := strings.Split(string(gomod), string('\n'))
	assert.Equal(t, gomodLines[0], "module testmod")
	assert.Len(t, gomodLines[1], 0)
	assert.Contains(t, gomodLines[2], "go 1.")
}

func TestWriteGoModToFile(t *testing.T) {
	gomod, err := createGoMod("testmod")
	assert.NoError(t, err)

	tmpDir := filepath.Join(
		os.TempDir(),
		fmt.Sprintf("gosprout_test_%s", randomString(6)),
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
		fmt.Sprintf("gosprout_test_%s", randomString(6)),
	)
	err := createMainGo(tmpDir)
	assert.NoError(t, err)
	fileInfo, err := os.Stat(filepath.Join(tmpDir, _mainGoFileName))
	assert.NoError(t, err)
	assert.False(t, fileInfo.IsDir())
	_ = os.RemoveAll(tmpDir)
}
