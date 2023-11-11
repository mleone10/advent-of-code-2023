package slice

func Map[S, R any](arr []S, f func(S) R) []R {
	res := []R{}
	for _, v := range arr {
		res = append(res, f(v))
	}
	return res
}

func Filter[S any](arr []S, f func(S) bool) []S {
	res := []S{}
	for _, v := range arr {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}
