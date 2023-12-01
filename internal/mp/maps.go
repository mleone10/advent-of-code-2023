package mp

// Keys returns a slice of all keys within `m`.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	ks := []K{}
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}

// Merge combines multiple maps into one, returning a new map containing all elements of those in `mps`.  Duplicate elements are overwritten in an indeterminate manner.
func Merge[M ~map[K]V, K comparable, V any](mps ...M) M {
	ret := M{}
	for _, m := range mps {
		for k, v := range m {
			ret[k] = v
		}
	}
	return ret
}
