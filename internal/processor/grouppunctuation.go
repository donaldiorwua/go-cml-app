package processor

import "strings"

func handleGroupPunc(words []string) []string {
	if len(words) == 0 {
		return words
	}

	// join words to handle spaces first
	text := strings.Join(words, " ")

	// handle multi-character punctuation first
	specials := []string{"...", "!?", "!!", "??"}
	for _, sp := range specials {
		text = strings.ReplaceAll(text, " "+sp, sp)
		text = strings.ReplaceAll(text, " ,", ",")
		text = strings.ReplaceAll(text, ",", ", ")

	}

	// split back into words
	return strings.Fields(text)
}
