// Package luhn contains the implementation of the Exercism's go exercise 'Luhn'
package luhn

import (
	"regexp"
	"strconv"
)

// Valid determines if the given number is valid per the Luhn formula
func Valid(input string) bool {
	// Remove spaces from input
	parsedInput := regexp.MustCompile(`[[:space:]|\p{Z}]`).ReplaceAllString(input, "")

	// check whether the number is valid
	if !isValidNumber(parsedInput) {
		return false
	}

	// declared required variables for the execution
	var (
		sum              int  // final sum of the calculation
		shouldDoubleNext bool // whether the next digit should be doubled or not
		currentDigit     int  // current digit during the loop
	)

	// iterate backward over the number
	for i := len(parsedInput) - 1; i >= 0; i-- {
		// convert to int
		currentDigit, _ = strconv.Atoi(string(parsedInput[i]))

		if shouldDoubleNext {
			// double the digit every second starting from right
			currentDigit = doubleDigit(currentDigit)
		}

		// add the result to the sum and reverse the condition to check the next
		sum += currentDigit
		shouldDoubleNext = !shouldDoubleNext
	}

	// a valid number if the sum is evenly divisible by 10
	return sum%10 == 0
}

// isValidNumber determines whether the input has more than 1 digit and is composed on digits only
func isValidNumber(input string) bool {
	return len(input) > 1 && regexp.MustCompile(`^[[:digit:]]*$`).MatchString(input)
}

// func doubleDigit multiples the input by 2 and substracts 9 if the result is greater or equal than 9
func doubleDigit(input int) (result int) {
	result = input * 2
	if result > 9 {
		result -= 9
	}
	return
}
