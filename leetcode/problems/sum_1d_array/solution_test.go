package sum_1d_array

import "testing"

var testCases = []struct {
	name     string
	input    []int
	expected []int
}{
	{
		name:     "Test with multiple elements",
		input:    []int{1, 2, 3, 4},
		expected: []int{1, 3, 6, 10},
	},
	{
		name:     "Test with single element",
		input:    []int{5},
		expected: []int{5},
	},
}

func TestRunningSum(t *testing.T) {

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := RunningSum(test.input)
			for i := range result {
				if result[i] != test.expected[i] {
					t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, result)
				}
			}
		})
	}
}

func TestRunningSumInPlance(t *testing.T) {

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := RunningSumInPlace(test.input)
			for i := range result {
				if result[i] != test.expected[i] {
					t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, result)
				}
			}
		})
	}
}
