// Package weather provides tools that
// forecast the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition stores the current weather condition("sunny", "rainy" etc)
// for the selected location in Goblinocus.
var CurrentCondition string

// CurrentLocation stores the name of the city in Goblinocus
// for which the forecast is being made.
var CurrentLocation string

// Forecast returns a string that describes the weather
// of the specified city in Goblinocus.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
