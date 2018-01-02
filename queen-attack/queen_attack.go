// Package queenattack contains the implementation of the Exercism's go exercise 'Queen Attack'
package queenattack

import (
	"errors"
	"math"
	"strconv"
)

// testVersion is the current version of the test
const testVersion = 2

// CanQueenAttack determines whether w can attack b
func CanQueenAttack(w, b string) (attack bool, err error) {
	// check if the positions are valid
	if w == b || !isPositionValid(w) || !isPositionValid(b) {
		err = errors.New("Invalid")
		return
	}

	// can attack if on same row, column or diagonals
	attack = w[0] == b[0] || w[1] == b[1] || isDiagonal(w, b)
	return
}

// isPositionValid determines if the given position is valid
func isPositionValid(input string) (result bool) {
	if len(input) == 2 {
		// convert the rank part from string to int
		rank, err := strconv.Atoi(string(input[1]))

		// check if both are within the allowed ranges
		result = input[0] <= 'h' && err == nil && rank >= 1 && rank <= 8
	}

	return
}

// isDiagonal checks if the two given positions share common left or right diagonals
func isDiagonal(p1 string, p2 string) bool {
	// for code clarity, assign each part to variables
	p1File := rune(p1[0])
	p1Rank, _ := strconv.Atoi(string(p1[1]))
	p2File := rune(p2[0])
	p2Rank, _ := strconv.Atoi(string(p2[1]))

	// check if the absolute values of the diff of files and ranks are equal which indicates a diagonal
	return math.Abs(float64(p2File-p1File)) == math.Abs(float64(p2Rank-p1Rank))
}
