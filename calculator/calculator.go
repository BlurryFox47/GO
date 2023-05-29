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
	var sum float64 = 0
	for _, input := range inputs {
		sum = Add(sum, input)
	}
	return sum
}
func SubtractMany(inputs ...float64) float64 {
	var res float64 = inputs[0]
	for i := 1; i < len(inputs); i++ {
		res = Subtract(res, inputs[i])
	}
	return res
}
func MuiltiplyMany(inputs ...float64) float64 {
	var sum float64 = 1
	for _, input := range inputs {
		sum = Multiply(sum, input)
	}
	return sum
}
func DivideMany(inputs ...float64) (float64, error) {
	var res float64 = inputs[0]
	var err error
	for i := 1; i < len(inputs); i++ {
		res, err = Divide(res, inputs[i])
		if err != nil {

			return 0, fmt.Errorf("One of dividers is a zero")
		}
	}
	return res, nil
}
