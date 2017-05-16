// Package bob contains the implementation of the Exercism's go exercise 'bob'
package bob

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// testVersion is the current version of the test
const testVersion = 2

// regex patterns used within the 'Hey' function
var (
	allSpacesPattern    = regexp.MustCompile(`^[[:space:]|\p{Z}]*$`)
	uselessCharsPattern = regexp.MustCompile(`\s|\d|,|!|:|\)`)
	allUpperCasePattern = regexp.MustCompile(`^[^a-z]*$`)
)

// Hey returns Bob's answer to the given input.
func Hey(input string) string {
	var result string

	if len(input) == 0 || allSpacesPattern.MatchString(input) {
		// if only spaces then 'Fine. Be that way!'
		result = "Fine. Be that way!"
	} else {
		// remove unnecessary chars and normalize input
		input = normalize(input)

		if allUpperCasePattern.MatchString(input) && len(input) > 1 {
			// if all alphabetic chars are upper case then 'Whoa, chill out!'
			result = "Whoa, chill out!"
		} else if strings.HasSuffix(input, "?") {
			// if ends in a question mark then 'Sure.'
			result = "Sure."
		} else {
			//  otherwise 'Whatever.'
			result = "Whatever."
		}
	}
	return result
}

// normalize normalizes the input converting chars
func normalize(input string) string {
	// function used during normalization
	isMn := func(r rune) bool {
		return unicode.Is(unicode.Mn, r)
	}

	// remove useless chars
	input = uselessCharsPattern.ReplaceAllString(input, "")

	// normalize input
	input, _, _ = transform.String(transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC), input)

	return input
}
