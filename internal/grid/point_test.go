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
