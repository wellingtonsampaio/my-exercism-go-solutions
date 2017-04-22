// Package gigasecond contains the implementation of the Exercism's go exercise 'gigasecond'.
package gigasecond

import "time"

// Constant declaration.
const testVersion = 4

// Represents a gigasecond. For performance boost, precalculate the gigasecond and reuse
const gigaSecondForDuration = time.Duration(1000000000) * time.Second

// AddGigasecond adds a gigasecond (10^9) to the given time.
func AddGigasecond(in time.Time) time.Time {
	return in.Add(gigaSecondForDuration)
}
