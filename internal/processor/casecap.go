package processor

import "strings"

func handleCap(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(cap)" && i > 0 {
			// capitalize the previous word
			word := words[i-1]
			if len(word) > 0 {
				words[i-1] = strings.Title(string(word))
			}

			// remove the "(cap)" token
			words = append(words[:i], words[i+1:]...)
			i-- // adjust index after removal
		}
	}
	return words
}
