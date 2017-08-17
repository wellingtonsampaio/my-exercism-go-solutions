// Package series contains the implementation of the Exercism's go exercise 'Series'
package series

// testVersion is the current version of the test
const testVersion = 2

// All returns a list of all substrings of s with length n.
func All(n int, s string) (result []string) {
	for start := 0; start <= len(s)-n; start++ {
		result = append(result, extract(n, s[start:start+n]))
	}
	return
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	return extract(n, s)
}

// First returns the first substring of s with length n or false if it is not possible
func First(n int, s string) (first string, ok bool) {
	defer func() {
		ok = recover() == nil
	}()

	first = UnsafeFirst(n, s)
	return
}

// extract returns a slice of the given string
func extract(n int, s string) string {
	return s[:n]
}
