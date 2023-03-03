package guf

// Ptr returns pointer to value.
func Ptr[T any](value T) *T {
	return &value
}

// PtrIf returns pointer to value if predicate is true.
func PtrIf[T any](value T, pred func() bool) *T {
	if !pred() {
		return nil
	}

	return &value
}

// DerefOrDefault returns dereferenced value of poiner or default value of type T.
func DerefOrDefault[T any](ptr *T) T {
	if ptr == nil {
		return *new(T)
	}

	return *ptr
}

// DerefOr returns dereferenced value of poiner or passed value.
func DerefOr[T any](ptr *T, defaultValue T) T {
	if ptr == nil {
		return defaultValue
	}

	return *ptr
}
