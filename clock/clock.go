// Package clock contains the implementation of the Exercism's go exercise 'clock'.
package clock

import (
	"fmt"
	"strconv"
)

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

// The number of minutes of a day
const dayInMinutes = 60 * 24

// Clock represents a clock where minutes can be added or substracted.
type Clock struct {
	minute int
}

// New creates a new value of type Clock.
func New(hour, minute int) Clock {
	// create a clock with 0 minutes and then add what was passed as arguments
	return Clock{0}.Add((hour * 60) + minute)
}

func (c Clock) String() string {
	// calculate the hours
	hour := c.minute / 60
	// process it until it is not greater than 23
	for hour > 23 {
		hour -= 24
	}

	// process the minutes
	minute := c.minute % 60

	// return the formatted time
	return fmt.Sprintf("%v:%v", addLeadingZero(hour), addLeadingZero(minute))
}

// Add manipulates the current time of the clock accepting positive and negative minutes.
func (c Clock) Add(minutes int) Clock {
	// sum the minutes
	c.minute += minutes

	// Convert negative time in the equivalent positive time
	for c.minute < 0 {
		c.minute += dayInMinutes
	}

	// ignore additional days
	for c.minute > dayInMinutes {
		c.minute -= dayInMinutes
	}

	return c
}

// addLeadingZero converts an int into string adding a leading 0 if the value is lesser than 10
func addLeadingZero(value int) string {
	result := strconv.Itoa(value)
	if value < 10 {
		result = "0" + result
	}
	return result
}
