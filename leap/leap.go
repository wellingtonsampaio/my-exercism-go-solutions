// Package leap contains the implementation of the Exercism's go exercise 'leap'.
package leap

import (
	"math"
)

const testVersion = 3

// IsLeapYear checks whether the given year is a leap year or not.
func IsLeapYear(year int) bool {

	// convert the year from int to float64
	yearAs64 := float64(year)

	// calculate the remainder of year / 4
	mod4 := math.Mod(yearAs64, 4)

	// calculate the remainder of year / 100
	mod100 := math.Mod(yearAs64, 100)

	// calculate the remainder of year / 400
	mod400 := math.Mod(yearAs64, 400)

	// if an even year divisible by 4 and not evenly divisible by 100 unless evenly divisible by 400 then is a leap year
	return (mod4 == 0 && mod100 > 0) || mod400 == 0
}
