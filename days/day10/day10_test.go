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

//go:embed test_input_3.txt
var testInput3 string

//go:embed input.txt
var input string

var p1Tcs = []struct {
	input           string
	expectedPartOne int
}{
	{
		testInput,
		4,
	},
	{
		testInput2,
		8,
	},
	{
		input,
		6856,
	},
}

var p2Tcs = []struct {
	input           string
	expectedPartTwo int
}{
	{
		testInput,
		1,
	},
	{
		testInput3,
		10,
	},
	{
		input,
		0,
	},
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range p1Tcs {
		p := day10.NewPipeField(tc.input)
		assert.Equals(t, tc.expectedPartOne, p.StepsFarthestFromStart())
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range p2Tcs {
		p := day10.NewPipeField(tc.input)
		assert.Equals(t, tc.expectedPartTwo, p.TilesEnclosedByLoop())
	}
}
