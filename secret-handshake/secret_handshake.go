// Package secret contains the implementation of the Exercism's go exercise 'Secret Handshake'
package secret

import (
	"strconv"
	"strings"
)

// testVersion is the current version of the test
const testVersion = 2

// Handshake accepts a decimal number and convert it to the appropriate sequence of events for a secret handshake.
func Handshake(code uint) (result []string) {

	// Convert the given decimal number into binary and split to an array
	binaryArray := strings.Split(strconv.FormatUint(uint64(code), 2), "")

	// Calculate the length of the array for later use
	binaryArrayLength := len(binaryArray)

	for pos := binaryArrayLength - 1; pos >= 0; pos-- {
		// Get the current part of the binary number and discard if starts with 0
		currentBinary := binaryArray[pos:binaryArrayLength]
		if currentBinary[0] == "0" {
			continue
		}

		// Convert to int
		number, _ := strconv.ParseInt(strings.Join(currentBinary, ""), 10, 32)

		// Append the secret codes to the handshake sequence
		if number == 1 {
			result = append(result, "wink")
		} else if number >= 10 && number < 100 {
			result = append(result, "double blink")
		} else if number >= 100 && number < 1000 {
			result = append(result, "close your eyes")
		} else if number >= 1000 && number < 10000 {
			result = append(result, "jump")
		} else if number > 10000 {
			// Reserse the array
			result = reverseArray(result)
		}
	}

	return
}

// reverseArray reserse the items of the given array
func reverseArray(array []string) (result []string) {
	for i := range array {
		n := array[len(array)-1-i]
		result = append(result, n)
	}
	return
}
