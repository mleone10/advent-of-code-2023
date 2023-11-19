package day20_test

import (
	"testing"

	day20 "github.com/mleone10/advent-of-code-2023/days/2015day20"
	"github.com/mleone10/advent-of-code-2023/internal/assert"
)

var tcs = []struct {
	input           int
	expectedPartOne int
	expectedPartTwo int
}{
	{
		10, 1, 0,
	},
	{
		70, 4, 0,
	},
	{
		150, 8, 0,
	},
	{
		36000000, 0, 0,
	},
}

func CalcPresents(t *testing.T) {
	ts := map[int]int{
		1: 10,
		4: 70,
		7: 80,
		8: 150,
	}
	for n, ps := range ts {
		assert.Equals(t, day20.CalcPresents(n), ps)
	}
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		actual := day20.FindHouseWithMinPresents(tc.input)
		assert.Equals(t, tc.expectedPartOne, actual)
	}
}

func TestSolvePartTwo(t *testing.T) {}
