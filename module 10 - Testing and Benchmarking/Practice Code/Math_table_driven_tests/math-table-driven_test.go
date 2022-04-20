package math

import "testing"

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
//this struct allow us to create a slice with multiple tests
//cases fot the function add where the arguments will be the inputs of the function
type addTest struct {
	arg1, arg2, expected int
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
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
