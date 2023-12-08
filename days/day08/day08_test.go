package day08_test

import (
	_ "embed"
	"testing"

	"github.com/mleone10/advent-of-code-2023/days/day08"
	"github.com/mleone10/advent-of-code-2023/internal/assert"
)

//go:embed test_input.txt
var testInput string

//go:embed test_input_2.txt
var testInput2 string

//go:embed input.txt
var input string

var tcs = []struct {
	input           string
	expectedPartOne int
	expectedPartTwo int
}{
	{
		testInput,
		2,
		0,
	},
	{
		testInput2,
		6,
		0,
	},
	{
		input,
		23147,
		0,
	},
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		m := day08.Map{Input: tc.input}
		assert.Equals(t, tc.expectedPartOne, m.ShortestTraversalDist())
	}
}

func TestSolvePartTwo(t *testing.T) {}
