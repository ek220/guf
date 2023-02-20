package f

// SliceToCh returns channel containing the slice values.
func SliceToCh[T any](s []T) <-chan T {
	ch := make(chan T)

	go func() {
		for _, v := range s {
			ch <- v
		}
		close(ch)
	}()

	return ch
}

// SliceFromCh returns slice containing the channel values.
func SliceFromCh[T any](ch <-chan T) []T {
	r := []T{}
	for v := range ch {
		r = append(r, v)
	}

	return r
}

// FilterCh returns a channel of the values of the source channel satisfying the predicate.
func FilterCh[T any](s <-chan T, pred func(T) bool) <-chan T {
	ch := make(chan T)

	go func() {
		for v := range s {
			if pred(v) {
				ch <- v
			}
		}
		close(ch)
	}()

	return ch
}

// MapCh returns a new channel from calling function for every source element.
func MapCh[T, R any](s <-chan T, f func(T) R) <-chan R {
	ch := make(chan R)

	go func() {
		for v := range s {
			ch <- f(v)
		}
		close(ch)
	}()

	return ch
}

// SplitCh devides source channel into two channel satisfying the predicate.
func SplitCh[T any](s <-chan T, pred func(T) bool) (<-chan T, <-chan T) {
	t := make(chan T)
	f := make(chan T)
	slice := SliceFromCh(s)

	go func(s <-chan T) {
		for v := range s {
			if pred(v) {
				t <- v
			}
		}
		close(t)
	}(SliceToCh(slice))

	go func(s <-chan T) {
		for v := range s {
			if !pred(v) {
				f <- v
			}
		}
		close(f)
	}(SliceToCh(slice))

	return t, f
}
