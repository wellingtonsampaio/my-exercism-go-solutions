// Package triangle contains the implementation of the Exercism's go exercise 'triangle'
package triangle

import (
	"math"
)

// testVersion is the current version of the test
const testVersion = 3

// Kind represents a type of triangule
type Kind string

const (
	// NaT indicates it is not a triangule
	NaT = Kind("NaT")
	// Equ represents an equilateral triangule
	Equ = Kind("Equ")
	// Iso represents an isosceles triangule
	Iso = Kind("Iso")
	// Sca represents a scalene triangule
	Sca = Kind("Sca")
	// Deg represents a degenerate triangule
	Deg = Kind("Deg")
)

// KindFromSides classify the triangule
func KindFromSides(a, b, c float64) Kind {
	// if a valid triangule
	if isValidTriangule(a, b, c) {

		if isEquilateral(a, b, c) {
			// Check if an equilateral one
			return Equ
		} else if isIsosceles(a, b, c) {
			// or isosceles
			return Iso
		} else if isScalene(a, b, c) {
			// or scalene
			return Sca
		} else if isDegenerate(a, b, c) {
			// or degenerate
			return Deg
		}
	}
	// otherwise not a triangule
	return NaT
}

// isValidTriangule determines if the given sides are a valid triangule
func isValidTriangule(a, b, c float64) bool {
	if !isValidSide(a) || !isValidSide(b) || !isValidSide(c) {
		return false
	} else if a+b < c || a+c < b || c+b < a {
		return false
	} else {
		return true
	}
}

// isValidSide determines if the given side is a valid one for a triangule
func isValidSide(side float64) bool {
	return !math.IsNaN(side) && !math.IsInf(side, 1) && side > 0
}

// isEquilateral checks if all sides of the triangule are equal
func isEquilateral(a, b, c float64) bool {
	return a == b && b == c
}

// isIsosceles checks if at least 2 sides of the triangule are equal
func isIsosceles(a, b, c float64) bool {
	return a == b || b == c || a == c
}

// isScalene checks if all sides of the triangule are different
func isScalene(a, b, c float64) bool {
	return a != b && b != c
}

// isDegenerate checks if the sum of 2 sides of the triangule are equal to third side
func isDegenerate(a, b, c float64) bool {
	return a+b == c || a+c == b || b+c == a
}
