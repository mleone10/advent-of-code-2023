package grid_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-2023/internal/assert"
	"github.com/mleone10/advent-of-code-2023/internal/grid"
)

func TestNeighbors(t *testing.T) {
	tcs := []struct {
		p                 grid.Point
		expectedNeighbors int
	}{
		{
			grid.Point{5, 10},
			8,
		},
		{
			grid.Point{0, 1},
			5,
		},
	}

	for _, tc := range tcs {
		actual := len(grid.Neighbors(tc.p))
		assert.Equals(t, tc.expectedNeighbors, actual)
	}
}

func TestAdd(t *testing.T) {
	subj := grid.Point{X: 5, Y: 5}
	actual := subj.Add(grid.Point{X: 2, Y: 4})

	assert.Equals(t, 7, actual.X)
	assert.Equals(t, 9, actual.Y)
}

func TestEquals(t *testing.T) {
	tcs := []struct {
		a, b     grid.Point
		expected bool
	}{
		{
			grid.Point{1, 2}, grid.Point{3, 4},
			false,
		},
		{
			grid.Point{4, 5}, grid.Point{4, 5},
			true,
		},
	}

	for _, tc := range tcs {
		assert.Equals(t, tc.expected, tc.a.Equals(tc.b))
	}
}
