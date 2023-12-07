package assert

import (
	"sort"
	"testing"

	"golang.org/x/exp/constraints"
)

// ArrayEquals tests whether two slices are equal as determined by their having identical lengths and, once sorted, identical elements at each index.
func ArrayEquals[T constraints.Ordered](t *testing.T, expected, actual []T) {
	if len(expected) != len(actual) {
		t.Errorf("expected %v with length [%v], got %v with length [%v]", expected, len(expected), actual, len(actual))
		return
	}

	var exp, act []T
	copy(expected, exp)
	copy(actual, act)

	sort.Slice(exp, func(i, j int) bool { return exp[i] < exp[j] })
	sort.Slice(act, func(i, j int) bool { return act[i] < act[j] })

	for i, v := range exp {
		if act[i] != v {
			t.Errorf("expected element [%v] of [%v] to be [%v], got [%v]", i, act, v, act[i])
			return
		}
	}
}

// Equals tests whether two comparable values are equal.
func Equals[T comparable](t *testing.T, expected, actual T) {
	if expected != actual {
		t.Errorf("expected [%v], got [%v]", expected, actual)
	}
}

// MapEquals tests whether two maps of comparable types are equal as determined by their having equal lengths, identical key sets, and identical values for each key in those sets.
func MapEquals[M ~map[K]V, K, V comparable](t *testing.T, expected, actual M) {
	if len(expected) != len(actual) {
		t.Errorf("expected length [%v], got [%v]", len(expected), len(actual))
	}

	for k, v := range expected {
		if actv, ok := actual[k]; ok {
			if actv != v {
				t.Errorf("expected value of [%v] to be [%v], got [%v]", k, v, actv)
			}
		} else {
			t.Errorf("key [%v] not found", k)
		}
	}
}
