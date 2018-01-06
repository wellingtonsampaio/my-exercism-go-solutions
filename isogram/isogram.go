// Package isogram contains the implementation of the Exercism's go exercise 'Isogram'
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram determines whether the given input is an isogram or not
func IsIsogram(input string) bool {
	// iterate over the letters of the input converting all to lower case
	for index, letter := range strings.ToLower(input) {
		if unicode.IsLetter(letter) {
			// if a letter, check if exists another one of the string
			if lastIndex := strings.LastIndexAny(input, string(letter)); lastIndex > index {
				// in case yes, then it is not an isogram
				return false
			}
		}
	}
	// otherwise the input is an isogram
	return true
}
