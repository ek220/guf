package guf

// Ptr return pointer to value.
func Ptr[T any](value T) *T {
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
