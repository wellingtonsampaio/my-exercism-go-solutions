// Package accumulate contains the implementation of the Exercism's go exercise 'accumulate'.
package accumulate

// testVersion is the current version of the test
const testVersion = 1

// Accumulate applies the converter to each element of the input slice and return the mapped slice
func Accumulate(input []string, converter func(string) string) []string {
	// if a nil or empty slice then return it
	if input == nil || len(input) == 0 {
		return input
	}

	// otherwise iterate over the elements of the slice and apply the converter function to each
	var mappedInput []string
	for _, element := range input {
		mappedInput = append(mappedInput, converter(element))
	}

	// return the mapped input
	return mappedInput
}
