package day09_test

import (
	_ "embed"
	"testing"

	"github.com/mleone10/advent-of-code-2023/days/day09"
	"github.com/mleone10/advent-of-code-2023/internal/assert"
	"github.com/mleone10/advent-of-code-2023/internal/slice"
)

//go:embed test_input.txt
var testInput string

//go:embed input.txt
var input string

var tcs = []struct {
	input           string
	expectedPartOne int
	expectedPartTwo int
}{
	{
		testInput,
		114,
		2,
	},
	{
		input,
		1666172641,
		933,
	},
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		actual := slice.Reduce(slice.TrimSplit(tc.input), 0, func(l string, ret int) int {
			return ret + day09.Next(l)
		})
		assert.Equals(t, tc.expectedPartOne, actual)
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		actual := slice.Reduce(slice.TrimSplit(tc.input), 0, func(l string, ret int) int {
			return ret + day09.Prev(l)
		})
		assert.Equals(t, tc.expectedPartTwo, actual)
	}
}
