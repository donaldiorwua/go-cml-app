package processor

import (
	"strconv"
	"strings"
)

func handleLowN(words []string, index int) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(low," && i+1 < len(words) {
			// extract the number from the next token, e.g., "2)"
			numStr := strings.TrimSuffix(words[i+1], ")")
			n, err := strconv.Atoi(numStr)
			if err != nil {
				continue
			}

			// lowercase the previous n words
			for j := 1; j <= n && i-j >= 0; j++ {
				words[i-j] = strings.ToLower(words[i-j])
			}

			// remove "(low," and "n)"
			words = append(words[:i], words[i+2:]...)
			i-- // adjust index after removal
		}
	}
	return words
}
