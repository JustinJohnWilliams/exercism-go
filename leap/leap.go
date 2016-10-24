package leap

const testVersion = 2

// IsLeapYear will return true or false if the year passed in is a leap year
func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}

	if year%4 == 0 && year%100 != 0 {
		return true
	}

	return false
}
