package secret

import "strconv"

const (
	wink        = "wink"
	doubleBlink = "double blink"
	closeEyes   = "close your eyes"
	jump        = "jump"
	testVersion = 1
)

// Handshake will take in a code and return the correct secret code
func Handshake(c uint) []string {
	code := strconv.FormatInt(int64(c), 2)
	length := len(code)

	if code == "0" || length == 0 {
		return nil
	}

	var phrase = []string(nil)

	if length > 0 && string(code[length-1]) == "1" {
		phrase = append(phrase, wink)
	}
	if length > 1 && string(code[length-2]) == "1" {
		phrase = append(phrase, doubleBlink)
	}
	if length > 2 && string(code[length-3]) == "1" {
		phrase = append(phrase, closeEyes)
	}
	if length > 3 && string(code[length-4]) == "1" {
		phrase = append(phrase, jump)
	}
	if length > 4 && string(code[length-5]) == "1" {
		phrase = reverse(phrase)
	}

	phrase = removeEmpties(phrase)

	if len(phrase) == 0 {
		return nil
	}

	return phrase
}

func removeEmpties(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func reverse(s []string) []string {
	var r []string
	for i := range s {
		n := s[len(s)-1-i]
		r = append(r, n)
	}

	return r
}
