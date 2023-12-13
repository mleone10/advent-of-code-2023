package day11_test

import (
	_ "embed"
	"testing"

	"github.com/mleone10/advent-of-code-2023/days/day11"
	"github.com/mleone10/advent-of-code-2023/internal/assert"
)

//go:embed test_input.txt
var testInput string

//go:embed input.txt
var input string

var tcs = []struct {
	input           string
	expansionFactor int
	expectedPartOne int
	expectedPartTwo int
}{
	{
		testInput,
		100,
		374,
		8410,
	},
	{
		input,
		1000000,
		9599070,
		842645913794,
	},
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		u := day11.NewUniverse(tc.input, 1)
		assert.Equals(t, tc.expectedPartOne, u.SumShortestPaths())
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		u := day11.NewUniverse(tc.input, tc.expansionFactor)
		assert.Equals(t, tc.expectedPartTwo, u.SumShortestPaths())
	}
}
