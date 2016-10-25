package diffsquares

import "math"

//SquareOfSums will square all the sums of the ints up to value
func SquareOfSums(value int) int {
	result := 0
	for i := 0; i <= value; i++ {
		result += i
	}
	return (result * result)
}

//SumOfSquares will sum all the squares of the ints up to the value
func SumOfSquares(value int) int {
	result := 0
	for i := 0; i <= value; i++ {
		result += (i * i)
	}

	return result
}

//Difference will find the difference between the Square of Sums and the Sum of Squares
func Difference(value int) int {
	return int(math.Abs(float64(SumOfSquares(value) - SquareOfSums(value))))
}
