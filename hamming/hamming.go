// Package hamming contains the implementation of the Exercism's go exercise 'hamming'.
package hamming

import (
	"errors"
	"strings"
)

const testVersion = 5

// Distance counts the number of differences between the given DNA strands
func Distance(a, b string) (int, error) {
	// default values for the case where strings match
	var (
		count int   // 0
		err   error // nil
	)

	// if the len of the strings are different, return error
	if len(a) != len(b) {
		count = -1
		err = errors.New("DNA strands are not of the same size")
	} else if a != b {
		// if they are not equal, check how many chars are different

		// split a and b into slices for performance improvement over accessing directly the index on the string
		splitA := strings.Split(a, "")
		splitB := strings.Split(b, "")

		for i := 0; i < len(splitA); i++ {
			// increment the counter when the chars at the same position are different
			if splitA[i] != splitB[i] {
				count++
			}
		}
	}

	// return the result according to the processing
	return count, err
}
