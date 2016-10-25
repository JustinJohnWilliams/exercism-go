package raindrops

import "strconv"

const testVersion = 2

// Convert will spit out the sound of the rain drop based off of it's facators
func Convert(rain int) string {
	var x = ""
	if rain%3 == 0 {
		x += "Pling"
	}

	if rain%5 == 0 {
		x += "Plang"
	}

	if rain%7 == 0 {
		x += "Plong"
	}

	if rain%3 != 0 && rain%5 != 0 && rain%7 != 0 {
		x = strconv.Itoa(rain)
	}

	return x
}
