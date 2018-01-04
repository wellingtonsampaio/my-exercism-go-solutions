// Package twofer contains the implementation of the Exercism's go exercise 'Two Fer'
package twofer

import "fmt"

// baseOutput is the base phrase that gets replaces according to the input
const baseOutput string = "One for %v, one for me."

// ShareWith determines the output of the 2-fer based on the input
func ShareWith(input string) string {
	value := input
	if len(input) == 0 {
		value = "you"
	}
	return fmt.Sprintf(baseOutput, value)
}
