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

//go:embed test_input_3.txt
var testInput3 string

//go:embed input.txt
var input string

var part1Tcs = []struct {
	input           string
	expectedPartOne int
}{
	{
		testInput,
		2,
	},
	{
		testInput2,
		6,
	},
	{
		input,
		23147,
	},
}

var part2Tcs = []struct {
	input           string
	expectedPartTwo int
}{
	{
		testInput3,
		6,
	},
	{
		input,
		22289513667691,
	},
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range part1Tcs {
		m := day08.Map{Input: tc.input}
		assert.Equals(t, tc.expectedPartOne, day08.ShortestTraversalDist(m, day08.StartAtAAA, day08.EndAtZZZ))
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range part2Tcs {
		m := day08.Map{Input: tc.input}
		assert.Equals(t, tc.expectedPartTwo, day08.ShortestTraversalDist(m, day08.StartAtAPrefix, day08.EndAtZSuffix))
	}
}
