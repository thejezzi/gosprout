package util

import (
	"errors"
	"math/rand"
	"os"
	"strings"

	static "github.com/thejezzi/gosprout"
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

type wordIndex = int

const (
	RandomProject wordIndex = iota
	RandomPath
)

func GetWordList(index wordIndex) ([]string, error) {
	lists := strings.Split(static.RandomNames, "\n\n")
	if len(lists) < 2 {
		return []string{}, errNoWordLists
	}
	list := strings.Split(lists[int(index)], "\n")

	cleaned := []string{}
	for _, word := range list {
		if word != "" {
			cleaned = append(cleaned, word)
		}
	}

	if len(cleaned) == 0 {
		return []string{}, errNoWordsAvailable
	}

	return cleaned, nil
}

func RandomSample(i wordIndex) (string, error) {
	list, err := GetWordList(i)
	if err != nil {
		return "", err
	}
	randomIndex := rand.Intn(len(list))
	return list[randomIndex], nil
}
