package processor

import (
	"strconv"
	"strings"
)

// handleCapN finds all "(cap, n)" instructions in the slice,
// capitalizes the previous n words, and removes the instruction tokens.
func handleCapN(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(cap," && i+1 < len(words) {

			// extract number from next token e.g. "2)"
			numStr := strings.TrimSuffix(words[i+1], ")")
			n, err := strconv.Atoi(numStr)
			if err != nil {
				continue
			}

			// apply capitalization to previous n words
			for j := 1; j <= n && i-j >= 0; j++ {
				word := words[i-j]
				if len(word) > 0 {
					words[i-j] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
				}
			}

			// remove "(cap," and "n)"
			words = append(words[:i], words[i+2:]...)
			i-- // adjust index after removal
		}
	}
	return words
}
