package set

type setVal comparable

// A Set[T] is a structure which holds unique values of type T.
type Set[T setVal] map[T]struct{}

// Add stores an element `v` in the Set.
func (s Set[T]) Add(vals ...T) {
	for _, v := range vals {
		s[v] = struct{}{}
	}
}

// Size returns the number of unique elements in the Set.
func (s Set[T]) Size() int {
	return len(s)
}
