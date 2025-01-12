package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsureEnv(t *testing.T) {
	randomEnvVar := RandomString(12)

	val, ok := os.LookupEnv(randomEnvVar)
	assert.False(t, ok)
	assert.Empty(t, val)

	wow := ensureEnv(randomEnvVar, "wow")
	val, ok = os.LookupEnv(randomEnvVar)
	assert.False(t, ok)
	assert.Empty(t, val)
	assert.Equal(t, wow, "wow")

	os.Setenv(randomEnvVar, "wow2")
	wow = ensureEnv(randomEnvVar, "defaultneverused")
	assert.Equal(t, wow, "wow2")
}

func TestDiscardEmptyElements(t *testing.T) {
	sample := []string{"eins", "", "zwei", "drei"}
	expected := []string{"eins", "zwei", "drei"}
	sample = discardEmptyElements(sample)
	assert.Equal(t, sample, expected)
}
