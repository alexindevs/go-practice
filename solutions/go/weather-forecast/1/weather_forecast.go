// Package weather provides tools for manipulating and reading weather data.
package weather

var (
    // CurrentCondition represents the city to be forecasted.
	CurrentCondition string
    // CurrentLocation represents the condition to be forecasted.
	CurrentLocation  string
)

// Forecast Function provides the forecast of a given city, and the condition in a formatted string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
