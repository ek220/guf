package f

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

func SliceFromCh[T any](ch <-chan T) []T {
	r := []T{}
	for v := range ch {
		r = append(r, v)
	}
	return r
}

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
