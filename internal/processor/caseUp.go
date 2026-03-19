package processor

import "strings"

func handleUp(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(up)" && i > 0 {
			// uppercase the previous word
			words[i-1] = strings.ToUpper(words[i-1])

			// remove the "(up)" token
			words = append(words[:i], words[i+1:]...)
			i-- // adjust index after removal
		}
	}
	return words
}
