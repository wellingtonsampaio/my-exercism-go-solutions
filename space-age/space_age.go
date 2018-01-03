// Package space contains the implementation of the Exercism's go exercise 'Space Age'
package space

// Planet struct represents a planet
type Planet string

// earthYearInSeconds number of seconds of one year on the Earth
const earthYearInSeconds float64 = 31557600

// Age given the age in seconds and the planet, determine the age of the person in Earth-years
func Age(seconds float64, planet Planet) float64 {
	// planets is a list of planets and their orbital period in Earth years
	planets := map[Planet]float64{
		"Earth":   1,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	return seconds / (earthYearInSeconds * planets[planet])
}
