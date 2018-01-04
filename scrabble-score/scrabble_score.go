// Package scrabble contains the implementation of the Exercism's go exercise 'Scrabble Score'
package scrabble

import "unicode"

// Score determines the score of the given input
func Score(input string) (result int) {

	// map of runes and their score
	scores := map[rune]int{
		'a': 1,
		'b': 3,
		'c': 3,
		'd': 2,
		'e': 1,
		'f': 4,
		'g': 2,
		'h': 4,
		'i': 1,
		'j': 8,
		'k': 5,
		'l': 1,
		'm': 3,
		'n': 1,
		'o': 1,
		'p': 3,
		'q': 10,
		'r': 1,
		's': 1,
		't': 1,
		'u': 1,
		'v': 4,
		'w': 4,
		'x': 8,
		'y': 4,
		'z': 10,
	}

	// iterate over the runes of the input
	for _, letter := range input {
		// if an upper case rune, convert to lower case
		if unicode.IsUpper(letter) {
			letter = unicode.ToLower(letter)
		}
		// add the rune's score to the result
		result += scores[letter]
	}
	return
}
