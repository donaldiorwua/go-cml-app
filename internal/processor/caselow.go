package processor

import "strings"

func handleLow(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(low)" && i > 0 {
			// lowercase the previous word
			words[i-1] = strings.ToLower(words[i-1])

			// remove the "(low)" token
			words = append(words[:i], words[i+1:]...)
			i-- // adjust index after removal
		}
	}
	return words
}
