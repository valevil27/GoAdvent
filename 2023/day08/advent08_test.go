package day08_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valevil27/adventofgo/2023/day08"
)

type Test struct {
	expected uint64
	input    string
	function func(string) uint64
}

func TestSolution(t *testing.T) {
	tests := []Test{
		{
			expected: 2,
			input:    "test-input",
			function: day08.Part1,
		},
		{
			expected: 6,
			input:    "test-input2",
			function: day08.Part1,
		},
		{
			expected: 6,
			input:    "test-input3",
			function: day08.Part2,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, test.function(test.input))
	}
}
