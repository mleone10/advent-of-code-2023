package grid_test

import (
	"strconv"
	"testing"

	"github.com/mleone10/advent-of-code-2023/internal/assert"
	"github.com/mleone10/advent-of-code-2023/internal/grid"
)

func TestGridMapIntToString(t *testing.T) {
	intGrid := grid.Grid[int]{}
	intGrid.Set(1, 1, 5)
	intGrid.Set(100, 100, 20)

	actual := grid.Map(intGrid, func(g grid.Grid[int], x, y int, v int) string {
		return strconv.Itoa(v)
	})

	v, ok := actual.Get(1, 1)
	assert.Equals(t, true, ok)
	assert.Equals(t, "5", v)

	v, ok = actual.Get(100, 100)
	assert.Equals(t, true, ok)
	assert.Equals(t, "20", v)
}

func TestGridMapStringToInt(t *testing.T) {
	stringGrid := grid.Grid[string]{}
	stringGrid.Set(1, 1, "five")
	stringGrid.Set(100, 100, "twenty")

	actual := grid.Map(stringGrid, func(g grid.Grid[string], x, y int, v string) int {
		return len(v)
	})

	v, ok := actual.Get(1, 1)
	assert.Equals(t, true, ok)
	assert.Equals(t, 4, v)

	v, ok = actual.Get(100, 100)
	assert.Equals(t, true, ok)
	assert.Equals(t, 6, v)
}

func TestFilter(t *testing.T) {
	g := grid.Grid[int]{}
	g.Set(1, 1, 2)
	g.Set(2, 2, 5)
	g.Set(3, 3, 6)
	g.Set(4, 4, 7)

	actual := grid.Filter(g, func(g grid.Grid[int], x, y int, v int) bool {
		return v%2 == 0
	})

	v, ok := actual.Get(1, 1)
	assert.Equals(t, true, ok)
	assert.Equals(t, 2, v)

	_, ok = actual.Get(2, 2)
	assert.Equals(t, false, ok)
}

func TestReduce(t *testing.T) {
	g := grid.Grid[int]{}
	g.Set(1, 1, 2)
	g.Set(2, 2, 4)
	g.Set(3, 3, 20)
	g.Set(4, 4, 6)

	actual := grid.Reduce(g, 0, func(g grid.Grid[int], x, y int, v int, init int) int {
		if v > init {
			return v
		}
		return init
	})

	assert.Equals(t, 20, actual)
}

func TestLength(t *testing.T) {
	g := grid.Grid[string]{}
	g.Set(1, 1, "foo")
	g.Set(2, 2, "bar")
	g.Set(3, 3, "fizz")
	g.Set(4, 4, "buzz")

	actual := grid.Length(g)

	assert.Equals(t, 4, actual)
}
