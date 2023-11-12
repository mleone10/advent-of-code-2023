package assert

import "testing"

func ArrayEquals[T comparable](t *testing.T, expected, actual []T) {
	if len(expected) != len(actual) {
		t.Errorf("expected [%v] with length [%v], got [%v] with length [%v]", expected, len(expected), actual, len(actual))
		return
	}
	for i, v := range expected {
		if actual[i] != v {
			t.Errorf("expected element [%v] of [%v] to be [%v], got [%v]", i, actual, v, actual[i])
			return
		}
	}
}

func Equals[T comparable](t *testing.T, expected, actual T) {
	if expected != actual {
		t.Errorf("expected [%v], got [%v]", expected, actual)
	}
}
