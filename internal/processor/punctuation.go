package processor

import "strings"

func handlePunctuation(words []string) []string {
	punctuations := ",.!?:;"

	for i := 0; i < len(words); i++ {
		word := words[i]

		// skip multi-character punctuation like ... or !?
		if len(word) > 1 {
			continue
		}

		if strings.Contains(punctuations, word) {
			// attach to previous word
			if i > 0 {
				words[i-1] += word
			}
			// remove the punctuation token
			words = append(words[:i], words[i+1:]...)
			i-- // adjust index after removal
		}
	}

	return words
}
