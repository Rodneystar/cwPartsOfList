package cwpartsoflist

import "strings"

//PartList : cw function
func PartList(words []string) string {
	var result string
	for i := 0; i < len(words)-1; i++ {
		result += getPairOfStrings(words, i)
	}
	return result
}

func getPairOfStrings(words []string, index int) string {
	var first, second []string
	for i := range words {
		if i <= index {
			first = append(first, words[i])
		} else {
			second = append(second, words[i])
		}
	}
	return "(" + strings.Join(first, " ") + ", " + strings.Join(second, " ") + ")"
}
