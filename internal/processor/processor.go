package processor

import "strings"

func Process(text string) string {

	word := strings.Fields(text)

	return strings.Join(word, " ")
}
