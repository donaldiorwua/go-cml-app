package processor

import (
	"strconv"
)

func Conversion(words []string) []string {

	for i := 0; i < len(words); i++ {

		//convert the string words from hex and binary to decimal.
		var base int
		if words[i] == "(hex)" && i > 0 {
			base = 16
		}
		if words[i] == "(bin)" && i > 0 {
			base = 2
		}

		if base != 0 {
			decimal, err := strconv.ParseInt(words[i-1], base, 64)
			if err == nil {
				words[i-1] = strconv.Itoa(int(decimal))
			}
			//Add the converted decimal word back into the slice
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return words
}
