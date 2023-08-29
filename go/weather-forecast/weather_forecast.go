// Package weather provides tools to forecast the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string
// CurrentLocation represents the current city being referred to.
var CurrentLocation string

// Forecast returns the weather condition for the given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
