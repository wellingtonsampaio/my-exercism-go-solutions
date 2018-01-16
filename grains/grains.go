// Package grains contains the implementation of the Exercism's go exercise 'Grains'
package grains

import "errors"

// Square calculates the number of grains given that the number on each square doubles
func Square(input int) (grains uint64, invalid error) {
	if input <= 0 || input > 64 {
		invalid = errors.New("Invalid input")
	} else {
		grains = 1 << uint(input-1)
	}
	return
}

// Total returns an expected number
func Total() uint64 {
	return 1<<64 - 1
}
