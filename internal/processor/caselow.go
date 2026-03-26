package processor

import "strings"

func handleLow(words []string) []string {
	for i := len(words) - 1; i >= 0; i-- {
		if words[i] == "(low)" && i > 0 {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}
