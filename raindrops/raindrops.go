// Package raindrops contains the implementation of the Exercism's go exercise 'raindrops'.
package raindrops

import (
	"bytes"
	"strconv"
)

// testVersion is the current version of the test
const testVersion = 2

// Convert convert a number to a string according to its factors
func Convert(number int) string {
	// for performance use bytes.Buffer to concatenate strings
	var result bytes.Buffer

	// If the number has 3 as a factor, add 'Pling'
	if number%3 == 0 {
		result.WriteString("Pling")
	}

	// If the number has 5 as a factor, add 'Plang'
	if number%5 == 0 {
		result.WriteString("Plang")
	}

	// If the number has 7 as a factor, add 'Plong'
	if number%7 == 0 {
		result.WriteString("Plong")
	}

	// Otherwise add the digits
	if result.Len() == 0 {
		result.WriteString(strconv.Itoa(number))
	}

	return result.String()
}
