// Package reverse contains the implementation of the Exercism's go exercise 'Reverse String'
package reverse

import "bytes"

// String reverse the chars of the given input
func String(input string) string {
	var result bytes.Buffer

	for i := len(input) - 1; i >= 0; i-- {
		result.WriteByte(input[i])
	}

	return result.String()
}
