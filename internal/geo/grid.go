package geo

type Grid[T comparable] struct {
	points map[int]map[int]T
}

// Set loads value v into the grid at location (x, y).  It handles all required initialization internally.
func (sg *Grid[T]) Set(x, y int, v T) {
	if sg.points == nil {
		sg.points = map[int]map[int]T{}
	}
	if _, ok := sg.points[y]; !ok {
		sg.points[y] = map[int]T{}
	}
	sg.points[y][x] = v
}

// Get retrieves the value in the grid at location (x, y).  If an item exists at that location, it returns it and a boolean set to true.  Otherwise, it returns the zero value of that item's type and a boolean set to false.
func (sg *Grid[T]) Get(x, y int) (T, bool) {
	if v, ok := sg.points[y][x]; ok {
		return v, true
	}
	return *new(T), false
}

func (sg *Grid[T]) GetPoint(p Point) (T, bool) {
	return sg.Get(p.X, p.Y)
}

// Map applies the given function `f` to all elements of the grid, returning a new grid to the caller.
func Map[T, V comparable](g Grid[T], f func(g Grid[T], x, y int, v T) V) Grid[V] {
	ng := Grid[V]{}
	for y, r := range g.points {
		for x, c := range r {
			ng.Set(x, y, f(g, x, y, c))
		}
	}
	return ng
}

// Filter applies the given function `f` to all elements of the grid, returning a new grid containing only those elements for which `f` returns true.
func Filter[T comparable](g Grid[T], f func(g Grid[T], x, y int, v T) bool) Grid[T] {
	ng := Grid[T]{}
	for y, r := range g.points {
		for x, c := range r {
			if f(g, x, y, c) {
				ng.Set(x, y, c)
			}
		}
	}
	return ng
}

// Reduce applies a function `f` to all elements of a grid and aggregates the result over subsequent calls of `f`.  The last result of `f` is returned.
func Reduce[T comparable, V any](g Grid[T], init V, f func(g Grid[T], x, y int, v T, res V) V) V {
	ret := init
	for y, r := range g.points {
		for x, c := range r {
			ret = f(g, x, y, c, ret)
		}
	}
	return ret
}

// Length returns the number of elements in the grid.  To count the number of elements matching a given condition, consider first applying a `grid.Filter`.
func Length[T comparable](g Grid[T]) int {
	return Reduce(g, 0, func(g Grid[T], x, y int, v T, init int) int {
		return init + 1
	})
}

// Points returns a slice of all populated grid.Points of the given grid.
func Points[T comparable](g Grid[T]) []Point {
	points := []Point{}
	for y, r := range g.points {
		for x := range r {
			points = append(points, Point{X: x, Y: y})
		}
	}
	return points
}
