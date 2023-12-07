package assert_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-2023/internal/assert"
)

func TestArrayEqualsInt(t *testing.T) {
	tcs := []struct {
		a, b          []int
		expectedEqual bool
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3},
			false,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{5, 4, 3, 2, 1},
			true,
		},
	}

	for _, tc := range tcs {
		u := testing.T{}
		assert.ArrayEquals(&u, tc.a, tc.b)
		if u.Failed() == tc.expectedEqual {
			t.Errorf("expected [%v] and [%v] equality to be [%v], got [%v]", tc.a, tc.b, tc.expectedEqual, !u.Failed())
		}
	}
}

func TestEqualsInt(t *testing.T) {
	tcs := []struct {
		a, b          int
		expectedEqual bool
	}{
		{
			5, 5, true,
		},
		{
			5, 10, false,
		},
	}

	for _, tc := range tcs {
		u := testing.T{}
		assert.Equals(&u, tc.a, tc.b)
		if u.Failed() == tc.expectedEqual {
			t.Errorf("expected [%v] and [%v] equality to be [%v], got [%v]", tc.a, tc.b, tc.expectedEqual, !u.Failed())
		}
	}
}
func TestEqualsString(t *testing.T) {
	tcs := []struct {
		a, b          string
		expectedEqual bool
	}{
		{
			"foo", "foo", true,
		},
		{
			"foo", "bar", false,
		},
	}

	for _, tc := range tcs {
		u := testing.T{}
		assert.Equals(&u, tc.a, tc.b)
		if u.Failed() == tc.expectedEqual {
			t.Errorf("expected [%v] and [%v] equality to be [%v], got [%v]", tc.a, tc.b, tc.expectedEqual, !u.Failed())
		}
	}
}

func TestMapEquals(t *testing.T) {
	tcs := []struct {
		a, b          map[string]int
		expectedEqual bool
	}{
		{
			map[string]int{"foo": 2, "bar": 4},
			map[string]int{"foo": 2, "bar": 4},
			true,
		},
		{
			map[string]int{"foo": 2, "bar": 4},
			map[string]int{"foo": 2},
			false,
		},
		{
			map[string]int{"foo": 2, "bar": 4},
			map[string]int{"foo": 2, "bar": 5},
			false,
		},
	}

	for _, tc := range tcs {
		u := testing.T{}
		assert.MapEquals(&u, tc.a, tc.b)
		if u.Failed() == tc.expectedEqual {
			t.Errorf("expected [%v] and [%v] equality to be [%v], got [%v]", tc.a, tc.b, tc.expectedEqual, !u.Failed())
		}
	}
}
