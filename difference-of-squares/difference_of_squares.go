// Package diffsquares contains the implementation of the Exercism's go exercise 'Difference Of Squares'
package diffsquares

import (
	"math"
)

// testVersion is the current version of the test
const testVersion = 1

// SquareOfSums sums the first N given natural numbers and return their square.
func SquareOfSums(n int) int {
	// sum the numbers
	sum := 0
	for n > 0 {
		sum += n
		n--
	}

	// return the square of the sum
	return int(math.Pow(float64(sum), 2))
}

// SumOfSquares sums the square of the first N given natural numbers.
func SumOfSquares(n int) int {
	if n > 0 {
		// recursively calculates the sum of the squares
		return int(math.Pow(float64(n), 2)) + SumOfSquares(n-1)
	}
	return n
}

// Difference returns the difference between the square of the sum and the sum of the squares of the first N given natural numbers.
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
