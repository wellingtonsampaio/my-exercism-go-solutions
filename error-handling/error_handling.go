// Package erratum contains the implementation of the Exercism's go exercise 'Error Handling'
package erratum

// testVersion is the current version of the test
const testVersion = 2

// Use sums the first N given natural numbers and return their square.
func Use(o ResourceOpener, input string) (result error) {
	// Attemp to open the resource and return if an error was found
	resource, err := open(o)
	if err != nil {
		// return the error received
		result = err
		return
	}
	// always close the resource
	defer resource.Close()

	// Handle panic condition and recover from it
	defer func() {
		// check if panic is in progress
		if recover := recover(); recover != nil {
			// if a FrobError, calls resource's defrob function
			if asFrob, ok := recover.(FrobError); ok {
				resource.Defrob(asFrob.defrobTag)
			}

			// if any error was thrown, return it
			if asError, ok := recover.(error); ok {
				result = asError
			}
		}
	}()

	// Call the frob function of the resource passing the input received
	resource.Frob(input)

	return
}

// open attempts to the open a resource. It may retry if a TransientError is returned.
func open(o ResourceOpener) (resource Resource, err error) {
	for {
		// Attempt to open the resouce
		resource, err = o()

		if _, ok := err.(TransientError); ok {
			// If a transient error, try again
			continue
		} else {
			// terminate if any other error occurred or the resource was opened successfully
			break
		}
	}
	return
}
