package math

import "testing"

// arg1 significa el argumento 1 and arg2 el argumento 2, and the expected stands for the 'el resultado que esperamos'
type addTest struct {
	arg1, arg2, expected int
}

type subtractTest struct {
	arg1, arg2, expected int
}

type factorialTest struct {
	arg1, expected int
}

var facrorialTests = []factorialTest{
	{0, 1},
	{10, 3628800},
	{7, 5040},
	{2, 2},
	{6, 720},
}

var subtractTests = []subtractTest{
	{2, 3, -1},
	{4, 8, -4},
	{6, 9, -3},
	{30, 10, 20},
}

var addTests = []addTest{
	{2, 3, 5},
	{4, 8, 12},
	{6, 9, 15},
	{3, 10, 13},
}

func TestAdd(t *testing.T) {

	for _, test := range addTests {
		if output := Add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func TestSubstract(t *testing.T) {

	for _, test := range subtractTests {
		if output := Subtract(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func TestFactorial(t *testing.T) {

	for _, test := range facrorialTests {
		if output := Factorial(test.arg1); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
