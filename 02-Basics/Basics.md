# Fahrenheit to Celsius
## Goals
- Learning the built in datatypes
- Perform basic arithmetic
- Intro to functions
- Intro to tests

## Variables and Datatypes
When writing code we many times want to save a value.  Sometimes we want to save it for use later, sometimes we want to just make reading our code easier.  All data on a computer at its base level is binary, but what those 1's and 0's represent can be anything.  Thanks to our awesome Go compiler we are given some base datatypes which are listed below.

### Concepts
#### Bits
Bits are the binary numbers that make up all datatypes.  Each bit represents 2 to the power of something.

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec eu arcu sagittis, imperdiet mi eu, facilisis turpis. Donec vitae imperdiet orci. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Duis elementum pretium risus, interdum malesuada lectus. Curabitur id feugiat est. Etiam maximus sem sit amet lorem hendrerit, nec ultrices nulla elementum. Sed in felis mi. 

#### Signed vs Unsigned
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec eu arcu sagittis, imperdiet mi eu, facilisis turpis. Donec vitae imperdiet orci. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Duis elementum pretium risus, interdum malesuada lectus. Curabitur id feugiat est. Etiam maximus sem sit amet lorem hendrerit, nec ultrices nulla elementum. Sed in felis mi. 

### Integer datatypes
These are all the datatypes that work with integers (whole numbers like 0, 1, 2, etc.)
|  Type | Description  | Uses  |
|---|---|---|
| `uint8`  | An 8 bit unsigned integer  | Mainly used when you want to store a small number ranging from 0 to 255 |
|   |   |   |
|   |   |   |

### Floating point datatypes
These are all the datatypes that work with floating point numbers (real numbers like 0.00234, 1.090, -122.345, etc.)
|  Type | Description  | Uses  |
|---|---|---|
| `bool`  | Short for boolean, this is used for basic logic.  It has 2 possible values `true` and `false`  | Mainly used when checking if something has happened or if you want to compare values  |
|   |   |   |
|   |   |   |

### Special datatypes
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec eu arcu sagittis, imperdiet mi eu, facilisis turpis. Donec vitae imperdiet orci. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Duis elementum pretium risus, interdum malesuada lectus. Curabitur id feugiat est. Etiam maximus sem sit amet lorem hendrerit, nec ultrices nulla elementum. Sed in felis mi. 

|  Type | Description  | Uses  |
|---|---|---|
| `bool`  | Short for boolean, this is used for basic logic.  It has 2 possible values `true` and `false`  | Mainly used when checking if something has happened or if you want to compare values  |
|   |   |   |
|   |   |   |

## Exercise
Implement `func FahrenheitToCelsius(temp float64) float64` and `func CelsiusToFahrenheit(temp float64) float64` found in `temperature.go`

`FahrenheitToCelsius` Requirements and Assumptions:
- Assume the input value is a temperature in Fahrenheit
- Convert and return the input value to Celsius

`CelsiusToFahrenheit` Requirements and Assumptions:
- Assume the input value is a temperature in Celsius
- Convert and return the input value to Fahrenheit

## Running your code
- Open a terminal (`cmd` or `powershell` on Windows)
- Change directories to the directory that contains `temperature.go` and `temperature_test.go`
- Run the command `go test`. You should see output that has PASS or FAIL on tests.