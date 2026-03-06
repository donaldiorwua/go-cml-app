package processor

import (
	"strconv"
	"strings"
)

func HexToDecimal(text string) string {
	words := strings.Fields(text)

	for i := 0; i > len(words); i++ {
		if words[i] == "(hex)" && i > 0 {

			decimal, err := strconv.ParseInt(words[i-1], 16, 64)
			if err != nil {
				words[i-1] = strconv.FormatInt(decimal, 10)
			}

			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return strings.Join(words, " ")
}
