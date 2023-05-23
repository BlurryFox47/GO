package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 2, b: 1, want: 3},
		{a: 1, b: 2, want: 3},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 2, want: 7},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%.3f,%.3f): want %.3f, got %.3f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 0},
		{a: 2, b: 1, want: 1},
		{a: 1, b: 2, want: -1},
		{a: 1, b: 1, want: 0},
		{a: 5, b: 2, want: 3},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Substract(%.3f,%.3f): want %.3f, got %.3f", tc.a, tc.b, tc.want, got)
		}
	}
}
func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 2, b: 1, want: 2},
		{a: 1, b: 2, want: 2},
		{a: 1, b: 1, want: 1},
		{a: 5, b: 2, want: 10},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%.3f,%.3f): want %.3f, got %.3f", tc.a, tc.b, tc.want, got)
		}
	}
}
func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b      float64
		want      float64
		errExpect bool
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 1, errExpect: false},
		{a: 6, b: 1, want: 6, errExpect: false},
		{a: 6, b: 2, want: 3, errExpect: false},
		{a: 1, b: 1, want: 1, errExpect: false},
		{a: 5, b: 2, want: 2.5, errExpect: false},
		{a: 3, b: 0, want: 0, errExpect: true},
	}
	for _, tc := range testCases {

		got, err := calculator.Divide(tc.a, tc.b)
		if tc.errExpect != (err != nil) {
			t.Fatalf("Division by zero")
		}
		if !tc.errExpect && !closeEnough(tc.want, got, 0.0001) {
			t.Errorf("Divide(%.3f,%.3f): want %.3f, got %.3f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a         float64
		want      float64
		errExpect bool
	}
	testCases := []testCase{
		{a: 9, want: 3, errExpect: false},
		{a: 4, want: 2, errExpect: false},
		{a: 25, want: 5, errExpect: false},
		{a: 81, want: 9, errExpect: false},
		{a: 2, want: 1.41421, errExpect: false},
		{a: -2, want: 0, errExpect: true},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		if tc.errExpect != (err != nil) {
			t.Fatalf("Negative argument")
		}
		if !closeEnough(tc.want, got, 0.0001) {
			t.Errorf("Sqrt(%.3f): want %.3f, gor %.3f", tc.a, tc.want, got)
		}
	}
}

func TestAddMany(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b, c, d, e, f, g float64
		want                float64
	}
	testCases := []testCase{
		{a: 5, b: 7, c: 9, d: 25, e: 2, f: 0, g: 0, want: 48},
	}
	for _, tc := range testCases {
		got := calculator.AddMany(tc.a, tc.b, tc.c, tc.d, tc.e, tc.f, tc.g)
		if got != tc.want {
			t.Errorf("AddMany ()")
		}
	}
}
