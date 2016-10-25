package acronym

import (
	"strings"
	"unicode"
)

const testVersion = 1

// Abbreviate will abbreviate the phrase passed in. e.g. Denial of Service = "DOS"
func Abbreviate(phrase string) string {
	result := ""
	words := strings.Split(phrase, " ")

	for _, word := range words {
		//split dashed words "mtel-oxide"
		if strings.ContainsAny(word, "-") {
			dashed := strings.Split(word, "-")
			for _, dash := range dashed {
				result += dash[:1]
			}
		} else if strings.ContainsAny(word, ":") {
			//get first initial of random acronyms "PHP: HyperText Processor"
			result += word[:1]
		} else {
			//take the word, and split it if there is any uppers "HyperText"
			result += word[:1]
			for _, letter := range word[len(word)-(len(word)-1):] {
				if unicode.IsUpper(letter) {
					result += string(letter)
				}
			}
		}
	}
	return strings.ToUpper(result)
}
