// Package weather a program that can forecast the current weather condition of various cities in Goblinocus.
package weather

var (
    // CurrentCondition  tell's the current condition. 
	CurrentCondition string
    // CurrentLocation tell's the current location.
	CurrentLocation  string
)

// Forecast it tells the weather of the city. 
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
