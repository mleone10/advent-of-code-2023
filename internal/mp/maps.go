package mp

func Keys[M map[K]V, K comparable, V any](m M) []K {
	ks := []K{}
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}
