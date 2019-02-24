package cwvowelcount

import "strings"

func VowelCount(str string) int {
	newSlice := Filter(strings.Split(str, ""),
		func(letter string) bool {
			return strings.ContainsAny(letter, "aeiou")
		})
	return len(newSlice)
}

func Filter(str []string, filterFn func(s string) bool) []string {
	newSlice := []string{}
	for _, char := range str {
		if filterFn(char) {
			newSlice = append(newSlice, char)
		}
	}
	return newSlice
}
