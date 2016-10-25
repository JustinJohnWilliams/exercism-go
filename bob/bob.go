package bob // package name must match the package name in bob_test.go
import (
	"fmt"
	"strings"
	"unicode"
)

const (
	fine        = "Fine. Be that way!"
	sure        = "Sure."
	whatever    = "Whatever."
	chill       = "Whoa, chill out!"
	testVersion = 2
)

// Hey is what bob says, the lazy bastard
func Hey(phrase string) string {
	fmt.Println(phrase)
	if phrase == "" || isSpace(phrase) {
		return fine
	}

	if isYelling(phrase) { //&& strings.Contains(phrase, "!") {
		return chill
	}

	if strings.Contains(phrase, "?") {
		return sure
	}

	return whatever
}

func isSpace(s string) bool {
	for _, character := range s {
		if !unicode.IsSpace(character) {
			return false
		}
	}
	return true
}

func isYelling(s string) bool {
	for _, character := range s {
		if unicode.IsLower(character) && string(character) != "!" && string(character) != "?" {
			return false
		}
	}
	return true
}
