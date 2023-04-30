package pointer

// From allocates a new variable of a given value and returns a pointer from it.
func From[T any](value T) *T {
	return &value
}
