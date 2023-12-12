package geo_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-2023/internal/assert"
	"github.com/mleone10/advent-of-code-2023/internal/geo"
)

func TestNeighbors(t *testing.T) {
	tcs := []struct {
		p                 geo.Point
		expectedNeighbors int
	}{
		{
			geo.Point{5, 10},
			8,
		},
		{
			geo.Point{0, 1},
			5,
		},
	}

	for _, tc := range tcs {
		actual := len(geo.Neighbors(tc.p))
		assert.Equals(t, tc.expectedNeighbors, actual)
	}
}

func TestAdd(t *testing.T) {
	subj := geo.Point{X: 5, Y: 5}
	actual := subj.Add(geo.Point{X: 2, Y: 4})

	assert.Equals(t, 7, actual.X)
	assert.Equals(t, 9, actual.Y)
}

func TestEquals(t *testing.T) {
	tcs := []struct {
		a, b     geo.Point
		expected bool
	}{
		{
			geo.Point{1, 2}, geo.Point{3, 4},
			false,
		},
		{
			geo.Point{4, 5}, geo.Point{4, 5},
			true,
		},
	}

	for _, tc := range tcs {
		assert.Equals(t, tc.expected, tc.a.Equals(tc.b))
	}
}

func TestTaxicabLength(t *testing.T) {
	assert.Equals(t, 6, geo.TaxicabLength(geo.Line{A: geo.Point{1, 1}, B: geo.Point{5, 3}}))
	assert.Equals(t, 0, geo.TaxicabLength(geo.Line{A: geo.Point{1, 1}, B: geo.Point{1, 1}}))
}
