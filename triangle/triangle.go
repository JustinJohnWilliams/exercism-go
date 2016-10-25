package triangle

import "math"

const testVersion = 3

// KindFromSides spits out the type of triangle it is from the given side lengths
func KindFromSides(a, b, c float64) Kind {
	if !isTriangle(a, b, c) {
		return NaT
	}

	if anyZeroes(a, b, c) || anyNaNs(a, b, c) || anyInfs(a, b, c) {
		return NaT
	}

	if isIso(a, b, c) {
		return Iso
	}

	if a == b && b == c {
		return Equ
	}

	return Sca
}

func isTriangle(a, b, c float64) bool {
	return ((a + b) >= c) && ((a + c) >= b) && ((b + c) >= a)
}

func anyZeroes(a, b, c float64) bool {
	return a <= 0 || b <= 0 || c <= 0
}

func anyNaNs(a, b, c float64) bool {
	return math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c)
}

func anyInfs(a, b, c float64) bool {
	return math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0)
}

func isIso(a, b, c float64) bool {
	if a == b && a != c {
		return true
	}

	if a == c && a != b {
		return true
	}

	if b == c && b != a {
		return true
	}

	return false
}

// Kind = Type of triangle
type Kind string

// Pick values for the following identifiers used by the test program.
const (
	NaT Kind = "Nat" // not a triangle
	Equ Kind = "Equ" // equilateral
	Iso Kind = "Iso" // isosceles
	Sca Kind = "Sca"
) // scalene

// Organize your code for readability.
