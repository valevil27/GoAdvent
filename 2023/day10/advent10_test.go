package day10_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valevil27/adventofgo/2023/day10"
)

type Test struct {
	expected int64
	input    string
	function func(string) int64
}

func TestSolution(t *testing.T) {
	tests := []Test{
		{
			expected: 4,
			input:    "test-input",
			function: day10.Part1,
		},
		{
			expected: 8,
			input:    "test-input2",
			function: day10.Part1,
		},
		{
			expected: 4,
			input:    "test-input3",
			function: day10.Part2,
		},
		{
			expected: 8,
			input:    "test-input4",
			function: day10.Part2,
		},
		{
			expected: 10,
			input:    "test-input5",
			function: day10.Part2,
		},
	}
	for i, test := range tests {
		assert.Equal(t, test.expected, test.function(test.input), fmt.Sprintf("Test #%v failed.\n", i+1))
	}
}
