package temperature

import (
	"testing"
	"math"
)

const (
	epsilon = 0.001
)

func areTemperatureEqualWithEpsilon(expected, returned float64) bool {
	// Since floating point numbers (real numbers) have a certian level of inaccuracy we need to compare with an epsilon value.  The epsilon value allows us to control the amount of error we deem allowed.
	return math.Abs(expected - returned) < epsilon
}

func compareExpectedAndReturnedTemp(expected, returned float64, t *testing.T) {
	if (!areTemperatureEqualWithEpsilon(expected, returned)) {
		t.Fatalf("Expected returned temperature to be within the range %v and %v.  Returned value was %v", expected - epsilon, expected + epsilon, returned)
	}
}

func TestFahrenheitToCelsius_ValidTemp(t *testing.T) {
	inputF := 50.0
	expectedC := 10.0
	returnedC := FahrenheitToCelsius(inputF)

	compareExpectedAndReturnedTemp(expectedC, returnedC, t)
}

func TestCelsiusToFahrenheit_ValidTemp(t *testing.T) {
	inputC := 123.0
	expectedF := 253.4
	returnedF := CelsiusToFahrenheit(inputC)

	compareExpectedAndReturnedTemp(expectedF, returnedF, t)
}

func TestFahrenheitToCelsius_NegativeTemp(t *testing.T) {
	inputF := -18.0
	expectedC := -27.7
	returnedC := FahrenheitToCelsius(inputF)

	compareExpectedAndReturnedTemp(expectedC, returnedC, t)
}

func TestCelsiusToFahrenheit_NegativeTemp(t *testing.T) {
	inputC := -17.0
	expectedF := 1.4
	returnedF := CelsiusToFahrenheit(inputC)

	compareExpectedAndReturnedTemp(expectedF, returnedF, t)
}

func TestFahrenheitToCelsius_ZeroTemp(t *testing.T) {
	inputF := 0.0
	expectedC := -17.7
	returnedC := FahrenheitToCelsius(inputF)

	compareExpectedAndReturnedTemp(expectedC, returnedC, t)
}

func TestCelsiusToFahrenheit_ZeroTemp(t *testing.T) {
	inputC := 0.0
	expectedF := 32.0
	returnedF := CelsiusToFahrenheit(inputC)

	compareExpectedAndReturnedTemp(expectedF, returnedF, t)
}

func TestFahrenheitToCelsius_BoilingTemp(t *testing.T) {
	inputF := 212.0
	expectedC := 100.0
	returnedC := FahrenheitToCelsius(inputF)

	compareExpectedAndReturnedTemp(expectedC, returnedC, t)
}

func TestCelsiusToFahrenheit_BoilingTemp(t *testing.T) {
	inputC := 100.0
	expectedF := 212.0
	returnedF := CelsiusToFahrenheit(inputC)

	compareExpectedAndReturnedTemp(expectedF, returnedF, t)
}

func TestFahrenheitToCelsius_ConvertBack(t *testing.T) {
	inputC := 79.0
	expectedC := inputC
	returnedC := FahrenheitToCelsius(CelsiusToFahrenheit(inputC))

	compareExpectedAndReturnedTemp(expectedC, returnedC, t)
}

func TestCelsiusToFahrenheit_ConvertBack(t *testing.T) {
	inputF := 92.0
	expectedF := inputF
	returnedF := CelsiusToFahrenheit(FahrenheitToCelsius(inputF))

	compareExpectedAndReturnedTemp(expectedF, returnedF, t)
}