package day10_test

import (
	_ "embed"
	"testing"

	"github.com/mleone10/advent-of-code-2023/days/day10"
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
		4,
		0,
	},
	{
		testInput2,
		8,
		0,
	},
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		p := day10.NewPipeField(tc.input)
		assert.Equals(t, tc.expectedPartOne, p.StepsFarthestFromStart())
	}
}

func TestSolvePartTwo(t *testing.T) {}
