// Package pangram contains the implementation of the Exercism's go exercise 'pangram'
package pangram

import (
	"strings"
	"unicode"
)

// testVersion is the current version of the test
const testVersion = 1

// IsPangram determines if the given input contains every letter of the alphabet at least once.
func IsPangram(input string) bool {
	// upper case latin letters
	var letters = unicode.Latin.R16[0]

	// convert input to uppercase for matching
	upperInput := strings.ToUpper(input)

	// iterate over the letters and return false if input does not contain any of them
	for c := letters.Lo; c <= letters.Hi; c += letters.Stride {
		if !strings.ContainsRune(upperInput, rune(c)) {
			return false
		}
	}
	// otherwise means input contains all the letters so return true
	return true
}
