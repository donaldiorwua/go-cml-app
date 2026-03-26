package processor

import (
	"strconv"
	"strings"
)

func handleLowN(words []string) []string {
	for i := len(words) - 1; i >= 0; i-- {
		if words[i] == "(low," && i+1 < len(words) {
			numStr := strings.TrimSuffix(words[i+1], ")")
			n, err := strconv.Atoi(numStr)
			if err != nil {
				continue
			}

			// lowercase the n words before i
			start := i - n
			if start < 0 {
				start = 0
			}
			for j := start; j < i; j++ {
				words[j] = strings.ToLower(words[j])
			}

			words = append(words[:i], words[i+2:]...)
		}
	}
	return words
}
