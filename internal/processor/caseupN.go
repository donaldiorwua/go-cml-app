package processor

import (
	"strconv"
	"strings"
)

func handleUpN(words []string) []string {
	for i := 0; i < len(words); i++ {

		// detect "(up," and next token like "2)"
		if words[i] == "(up," && i+1 < len(words) {

			// extract number from "2)"
			numStr := strings.TrimSuffix(words[i+1], ")")
			n, err := strconv.Atoi(numStr)
			if err != nil {
				continue
			}

			// apply uppercase to previous n words
			for j := 1; j <= n && i-j >= 0; j++ {
				words[i-j] = strings.ToUpper(words[i-j])
			}

			// remove "(up," and "2)"
			words = append(words[:i], words[i+2:]...)
			i-- // adjust index after removal
		}
	}
	return words
}
