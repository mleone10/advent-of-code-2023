package day05_test

import (
	_ "embed"
	"testing"

	"github.com/mleone10/advent-of-code-2023/days/day05"
	"github.com/mleone10/advent-of-code-2023/internal/assert"
)

//go:embed test_input.txt
var testInput string

//go:embed input.txt
var input string

var tcs = []struct {
	input           string
	seedCount       int
	expectedPartOne int
	expectedPartTwo int
}{
	{
		testInput,
		4,
		35,
		46,
	},
	{
		input,
		20,
		178159714,
		0,
	},
}

func TestNewAlmanac(t *testing.T) {
	for _, tc := range tcs {
		a := day05.NewAlmanac(tc.input)
		assert.Equals(t, tc.seedCount, len(a.Seeds))
	}
}

func TestRangeFunc(t *testing.T) {
	ranges := []string{"50 98 2", "52 50 48"}

	f := day05.NewRangeFunc(ranges)

	assert.Equals(t, 0, f(0))
	assert.Equals(t, 48, f(48))
	assert.Equals(t, 52, f(50))
	assert.Equals(t, 99, f(97))
	assert.Equals(t, 51, f(99))
	assert.Equals(t, 14, f(14))
	assert.Equals(t, 57, f(55))
}

func TestLocation(t *testing.T) {
	a := day05.NewAlmanac(tcs[0].input)

	assert.Equals(t, 82, a.Location(79))
	assert.Equals(t, 43, a.Location(14))
	assert.Equals(t, 86, a.Location(55))
	assert.Equals(t, 35, a.Location(13))
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		a := day05.NewAlmanac(tc.input)
		minLoc := a.ClosetLocation()
		assert.Equals(t, tc.expectedPartOne, minLoc)
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		a := day05.NewAlmanac(tc.input)
		minLoc := a.ClosetLocationSeedRange()
		assert.Equals(t, tc.expectedPartTwo, minLoc)
	}
}
