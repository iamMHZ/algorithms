package richest_customer_wealth

import "testing"

var testCases = []struct {
	name     string
	input    [][]int
	expected int
}{
	{

		name:     "find wealthiest when having customers of same wealth.",
		input:    [][]int{{1, 2, 3}, {3, 2, 1}},
		expected: 6,
	},
	{

		name:     "find wealthiest when having only one wealthy customer.",
		input:    [][]int{{1, 5}, {7, 3}, {3, 5}},
		expected: 10,
	},
}

func TestRunner(t *testing.T) {

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := MaximumWealth(test.input)
			if result != test.expected {
				t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, result)

			}
		})
	}
}
