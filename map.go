package guf

// CloneMap returns copy of source map.
func CloneMap[K comparable, V any](s map[K]V) map[K]V {
	r := make(map[K]V, len(s))
	for k, v := range s {
		r[k] = v
	}

	return r
}

// FilterMap returns a map of the values of the source map satisfying the predicate.
func FilterMap[K comparable, V any](s map[K]V, pred func(K, V) bool) map[K]V {
	r := map[K]V{}

	for k, v := range s {
		if pred(k, v) {
			r[k] = v
		}
	}

	return r
}

// MapMap returns a new map from calling function for every source key and value.
func MapMap[K, K1 comparable, V, V1 any](s map[K]V, f func(K, V) (K1, V1)) map[K1]V1 {
	r := make(map[K1]V1, len(s))

	for k, v := range s {
		k1, v1 := f(k, v)
		r[k1] = v1
	}

	return r
}

// SplitMap devides source map into two maps satisfying the predicate.
func SplitMap[K comparable, V any](s map[K]V, pred func(K, V) bool) (map[K]V, map[K]V) {
	t := map[K]V{}
	f := map[K]V{}

	for k, v := range s {
		if pred(k, v) {
			t[k] = v
		} else {
			f[k] = v
		}
	}

	return t, f
}
