package day07_test

import (
	_ "embed"
	"testing"

	"github.com/mleone10/advent-of-code-2023/days/day07"
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
		6592,
		6839,
	},
	{
		input,
		248453531,
		248781813,
	},
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		actual := day07.TotalWinnings(slice.TrimSplit(tc.input))
		assert.Equals(t, tc.expectedPartOne, actual)
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		day07.JokersWild()
		actual := day07.TotalWinnings(slice.TrimSplit(tc.input))
		assert.Equals(t, tc.expectedPartTwo, actual)
	}
}
