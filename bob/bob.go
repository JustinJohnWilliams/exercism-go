package bob // package name must match the package name in bob_test.go
import (
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
	// fmt.Println(phrase)
	if isSayingNothing(phrase) {
		return fine
	}

	if isYelling(phrase) {
		return chill
	}

	if isQuestion(phrase) {
		return sure
	}

	return whatever
}

func isSayingNothing(s string) bool {
	for _, character := range s {
		if !unicode.IsSpace(character) {
			return false
		}
	}
	return true
}

func isYelling(s string) bool {
	if !hasLetters(s) {
		return false
	}
	for _, character := range s {
		if unicode.IsLower(character) && string(character) != "!" && string(character) != "?" {
			return false
		}
	}
	return true
}

func isQuestion(s string) bool {
	return strings.HasSuffix(strings.TrimSpace(s), "?")
}

func hasLetters(s string) bool {
	return strings.ContainsAny(s, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
}
