package guf

// CloneSlice returns copy of source slice.
func CloneSlice[T any](s []T) []T {
	r := make([]T, len(s))
	copy(r, s)

	return r
}

// FilterSlice returns a slice of the values of the source slice satisfying the predicate.
func FilterSlice[T any](s []T, pred func(T) bool) []T {
	r := []T{}

	for _, v := range s {
		if pred(v) {
			r = append(r, v)
		}
	}

	return r
}

// MapSlice returns a new slice from calling function for every source element.
func MapSlice[T, R any](s []T, f func(T) R) []R {
	r := make([]R, len(s))
	for i, v := range s {
		r[i] = f(v)
	}

	return r
}

// SplitSlice devides source slice into two slices satisfying the predicate.
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

// ReverseSlice reverses slice.
func ReverseSlice[T any](s []T) []T {
	r := make([]T, len(s))
	for i := 0; i < len(s); i++ {
		r[i] = s[len(s)-1-i]
	}

	return r
}
