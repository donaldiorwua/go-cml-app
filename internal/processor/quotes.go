package processor

func handleQuotes(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "'" {
			j := i + 1
			for j < len(words) && words[j] != "'" {
				j++
			}
			if j > len(words) {
				break
			}
			if i+1 < len(words) {
				words[i+1] = "'" + words[i+1]
			}
			if j-1 > +0 {
				words[j-1] = words[j-1] + "'"
			}
			words = append(words[:j], words[j+1:]...)

			words = append(words[:i], words[i+1:]...)
			i = -1
		}
	}
	return words
}
