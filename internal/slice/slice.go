package slice

import "strings"

type summable interface {
	int
}

// Map applies a function `f` to each element in a slice `arr`, returning a new slice of the results of `f`.
func Map[S, R any](arr []S, f func(S) R) []R {
	res := []R{}
	for _, v := range arr {
		res = append(res, f(v))
	}
	return res
}

// Map applies a function `f` to each element in a slice `arr`, returning a new slice containing only those elements of `arr` for which `f` evaluates to true.
func Filter[S any](arr []S, f func(S) bool) []S {
	res := []S{}
	for _, v := range arr {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

// Reduce iterates through a slice `arr`, accumulating the result of `f` to some initial value `init`, then returning the final result.
func Reduce[R, S any](arr []S, init R, f func(s S, r R) R) R {
	res := init
	for _, s := range arr {
		res = f(s, res)
	}
	return res
}

// TrimSplit splits a string on newline characters after trimming leading and trailing whitespace.  This ensures that the resultant slice contains no empty strings.
func TrimSplit(s string) []string {
	return strings.Split(strings.TrimSpace(s), "\n")
}

// Sum is a convenience method to add up a slice of numbers (as defined by an internal type).
func Sum[V summable](ns []V) V {
	var ret V
	return Reduce(ns, ret, func(n, ret V) V {
		return ret + n
	})
}

// Contains returns true if slice `arr` contains a value equal to `v`.
func Contains[S comparable](arr []S, v S) bool {
	for _, a := range arr {
		if a == v {
			return true
		}
	}
	return false
}
