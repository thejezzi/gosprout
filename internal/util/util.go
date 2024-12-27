package util

import "os"

func EnsureEnv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}

func DiscardEmptyElements(oldSlice []string) []string {
	newSlice := []string{}
	for _, elem := range oldSlice {
		if len(elem) > 0 {
			newSlice = append(newSlice, elem)
		}
	}
	return newSlice
}
