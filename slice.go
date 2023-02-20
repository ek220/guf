package f

func CloneSlice[T any](s []T) []T {
	r := make([]T, len(s))
	for i, v := range s {
		r[i] = v
	}
	return r
}

func FilterSlice[T any](s []T, pred func(T) bool) []T {
	r := []T{}
	for _, v := range s {
		if pred(v) {
			r = append(r, v)
		}
	}
	return r
}

func MapSlice[T, R any](s []T, f func(T) R) []R {
	r := make([]R, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func SplitSlice[T any](s []T, pred func(T) bool) ([]T, []T) {
	t := []T{}
	f := []T{}
	for _, v := range s {
		if pred(v) {
			t = append(t, v)
		} else {
			f = append(f, v)
		}
	}
	return t, f
}
