package day01_test

import (
	_ "embed"
	"testing"

	"github.com/mleone10/advent-of-code-2023/days/day01"
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
		243,
		232,
	},
	{
		input,
		55123,
		55260,
	},
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		actual := day01.CalibrationSum(slice.TrimSplit(tc.input), day01.Numerics)
		assert.Equals(t, tc.expectedPartOne, actual)
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		actual := day01.CalibrationSum(slice.TrimSplit(tc.input), day01.Alphanumerics)
		assert.Equals(t, tc.expectedPartTwo, actual)
	}
}
