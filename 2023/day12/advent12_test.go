package day12_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valevil27/adventofgo/2023/day12"
)

type Test struct {
	expected int64
	input    string
	function func(string) int64
}

func TestSolution(t *testing.T) {
	tests := []Test{
		{
			expected: 21,
			input:    "test-input",
			function: day12.Part1,
		},
	}
	for i, test := range tests {
		assert.Equal(t, test.expected, test.function(test.input), fmt.Sprintf("Test #%v failed.\n", i+1))
	}
}
