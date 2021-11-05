//Package weather forecasts the weather in different cities.
package weather

//CurrentCondition is a string representing the current weather condition.
var CurrentCondition string

//CurrentLocations is a string representing the current city.
var CurrentLocation string

//Forecast is a function that takes two arguments - city and condition and returns a string representing the forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
