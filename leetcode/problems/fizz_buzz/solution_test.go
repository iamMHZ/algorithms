package fizz_buzz

import "testing"

var testCases = []struct {
	name     string
	input    int
	expected []string
}{
	{
		name:     "test with only having fizz",
		input:    3,
		expected: []string{"1", "2", "Fizz"},
	},
	{
		name:     "test with having both fizz and buzz",
		input:    5,
		expected: []string{"1", "2", "Fizz", "4", "Buzz"},
	},
	{
		name:  "test with having multiple fizz and buzz",
		input: 15,
		expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7",
			"8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz",
		},
	},
}

func TestRunnerFizzBuzz(t *testing.T) {

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := FizzBuzz(test.input)
			for i := range result {
				if result[i] != test.expected[i] {
					t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, result)
				}
			}
		})
	}
}
