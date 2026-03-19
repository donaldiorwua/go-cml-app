package processor

import "strings"

func Process(text string) string {

	//text = normalizePunctuation(text)

	lines := strings.Split(text, "\n")

	for i := range lines {
		words := strings.Fields(lines[i])

		words = Conversion(words)
		words = handleUpN(words)
		words = handleCap(words)
		words = handleLow(words)
		words = handleUp(words)
		words = handleCapN(words)
		words = handleLowN(words)
		words = handlePunctuation(words)
		words = handleGroupPunc(words)
		words = handleQuotes(words)
		words = handleGrammar(words)

		lines[i] = strings.Join(words, " ")
	}
	result := strings.Join(lines, "\n")
	return result
}

// func normalizePunctuation(text string) string {

// 	punct := []string{".", ",", "!", "?", ":", ";"}

// 	for _, p := range punct {
// 		text = strings.ReplaceAll(text, p, " "+p+" ")
// 	}

// 	return text
//}
