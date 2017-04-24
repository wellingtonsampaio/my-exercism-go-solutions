// Package acronym contains the implementation of the Exercism's go exercise 'acronym'
package acronym

import (
	"bytes"
	"regexp"
	"strings"
)

// testVersion is the current version of the test
const testVersion = 2

// Abbreviate returns the acronym of the given input phrase
func Abbreviate(input string) string {
	var result bytes.Buffer

	// split the input into tokens (maybe yperText could be handled differently)
	fields := regexp.MustCompile("\\s|-|yperTex").Split(input, -1)
	for _, field := range fields {
		// and get the first char of each token
		result.WriteString(string(field[0]))
	}

	// return converting the acronym to uppercase
	return strings.ToUpper(result.String())
}
