package processor

import (
	"strconv"
	"strings"
)

func handleCapN(words []string) []string {
	for i := 0; i < len(words); i++ {

		if words[i] == "(cap," && i+1 < len(words) {

			// extract number from "2)"
			numStr := strings.TrimSuffix(words[i+1], ")")
			n, err := strconv.Atoi(numStr)
			if err != nil {
				continue
			}

			// apply caps to previous n words
			for j := 1; j <= n && i-j >= 0; j++ {
				words[i-j] = strings.Title(words[i-j])
			}

			// remove "(up," and "2)"
			words = append(words[:i], words[i+2:]...)
			i-- // adjust index after removal
		}
	}
	return words
}
