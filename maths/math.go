package maths

import (
	"math"
)

func Round(number float64, decimalPlaces uint) float64 {
	ratio := math.Pow(10, float64(decimalPlaces))
	return math.Round(number*ratio) / ratio
}

// returns the absolute value as a float
func AbsoluteFloat(number float64) float64 {
	return math.Abs(number)
}

// returns the absolute value as an integer
func AbsoluteInt(number float64) int {
	number = math.Abs(number)
	intNumber := int(number)
	return intNumber
}

func ConvertCurrency(decimal float64, fromCurrency string, toCurrency string) float64 {
	var converted float64
	return converted

}

// returns the largest value from the list. Can be an int list or float list
// func LargestGetFromList(list []any) any {
// 	// before starting the loop, get the list type
// 	v := reflect.TypeOf(list)
// 	if strings.Contains(v.Name(), "float") {
// 		floatList := list
// 	} else if strings.Contains(v.Name(), "int") {
// 		intList := list
// 	}

// }
