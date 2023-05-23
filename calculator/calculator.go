package calculator

import (
	"fmt"
	"math"
)

func Add(a, b float64) float64 {
	return a + b
}
func Subtract(a, b float64) float64 {
	return a - b
}
func Multiply(a, b float64) float64 {
	return a * b
}
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("bad imput: %.3f,%.3f (division by zero is undefined)", a, b)
	}
	return a / b, nil
}
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("bad imput: %.3f cannot be negative", a)
	}
	return math.Sqrt(a), nil
}
func AddMany(inputs ...float64) float64 {
	// v debug-u mi prepisuje hodnoty sum a tmp na 0, tak mi to cyklicky nefunguje
	// použil som forr ale rozmýšlam že použijem klasický fori
	var sum, tmp float64 = 0, 0
	for _, input := range inputs {
		sum := Add(tmp, input)
		tmp := sum
		println("tmp %.3f", tmp)
	}
	return sum
}
