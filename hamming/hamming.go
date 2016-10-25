package hamming

import "errors"

const testVersion = 5

// Distance measures the Hamming distnace between two strands of DNA
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("string length must be the same")
	}

	dist := 0

	for i := range a {
		if a[i] != b[i] {
			dist++
		}
	}

	return dist, nil
}
