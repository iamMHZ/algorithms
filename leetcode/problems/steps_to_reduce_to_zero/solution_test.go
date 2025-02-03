package stepstoreducetozero

import "testing"

var testCases = []struct {
	name     string
	input    int
	expected int
}{

	{
		name:     "test with large number",
		input:    123,
		expected: 12,
	},
	{
		name:     "test with even number",
		input:    14,
		expected: 6,
	},
}

func TestRunnerFizzBuzz(t *testing.T) {

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := numberOfSteps(test.input)

			if result != test.expected {
				t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, result)
			}

		})
	}
}
