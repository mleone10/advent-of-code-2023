package assert

import (
	"sort"
	"testing"

	"golang.org/x/exp/constraints"
)

func ArrayEquals[T constraints.Ordered](t *testing.T, expected, actual []T) {
	if len(expected) != len(actual) {
		t.Errorf("expected [%v] with length [%v], got [%v] with length [%v]", expected, len(expected), actual, len(actual))
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

func Equals[T comparable](t *testing.T, expected, actual T) {
	if expected != actual {
		t.Errorf("expected [%v], got [%v]", expected, actual)
	}
}
