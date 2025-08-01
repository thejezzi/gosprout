package util

import (
	"errors"
	"math/rand"
	"os"
)

var (
	errNoWordLists      = errors.New("there is no word list to choose from")
	errNoWordsAvailable = errors.New("no words in random word list")
)

func ensureEnv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}

func discardEmptyElements(oldSlice []string) []string {
	newSlice := []string{}
	for _, elem := range oldSlice {
		if len(elem) > 0 {
			newSlice = append(newSlice, elem)
		}
	}
	return newSlice
}

func RandomString(length int) string {
	result := make([]byte, length)

	for i := range length {
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
