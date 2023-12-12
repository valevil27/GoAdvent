package day10_test

import (
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
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, test.function(test.input))
	}
}
