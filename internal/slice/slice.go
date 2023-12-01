package slice

import "strings"

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

// TrimSplit splits a string on newline characters after trimming leading and trailing whitespace.  This ensures that the resultant slice contains no empty strings.
func TrimSplit(s string) []string {
	return strings.Split(strings.TrimSpace(s), "\n")
}
