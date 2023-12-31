package day09_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valevil27/adventofgo/2023/day09"
)

type Test struct {
	expected int64
	input    string
	function func(string) int64
}

func TestSolution(t *testing.T) {
	tests := []Test{
		{
			expected: 114,
			input:    "test-input",
			function: day09.Part1,
		},
		{
			expected: 2,
			input:    "test-input",
			function: day09.Part2,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, test.function(test.input))
	}
}
