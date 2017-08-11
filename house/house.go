// Package house contains the implementation of the Exercism's go exercise 'House'
package house

import "bytes"

// testVersion is the current version of the test
const testVersion = 1

// variable representing the song in verses
var structure = [][]string{
	{"house that Jack built.", ""},
	{"malt", "lay in"},
	{"rat", "ate"},
	{"cat", "killed"},
	{"dog", "worried"},
	{"cow with the crumpled horn", "tossed"},
	{"maiden all forlorn", "milked"},
	{"man all tattered and torn", "kissed"},
	{"priest all shaven and shorn", "married"},
	{"rooster that crowed in the morn", "woke"},
	{"farmer sowing his corn", "kept"},
	{"horse and the hound and the horn", "belonged to"},
}

// the number of verses in the song
var numberOfVerses = len(structure)

// Song return the entire song
func Song() string {
	var result bytes.Buffer

	for index := 1; index <= numberOfVerses; index++ {
		// append each verse to the song
		result.WriteString(Verse(index))
		if !isLastVerse(index) {
			result.WriteString("\n\n")
		}
	}
	return result.String()
}

// Verse return the given n verse of the song
func Verse(index int) string {
	return buildVerse(index, true)
}

// buildVerse creates the requested verse recursively
func buildVerse(index int, includeBeginning bool) string {
	var result bytes.Buffer

	if index > 0 && index <= numberOfVerses {
		// get the current verse structure
		pattern := structure[index-1]

		// Include the beginning if requested
		if includeBeginning {
			result.WriteString("This is the ")
		}
		// Add the first part of the verse
		result.WriteString(pattern[0])

		// Add the second part of the verse and recursively add the previous verses
		if pattern[1] != "" {
			result.WriteString("\nthat ")
			result.WriteString(pattern[1])
			result.WriteString(" the ")
			result.WriteString(buildVerse(index-1, false))
		}
	}
	return result.String()
}

// isLastVerse determines whether the given index refers to the last verse of the song
func isLastVerse(index int) bool {
	return index == numberOfVerses
}
