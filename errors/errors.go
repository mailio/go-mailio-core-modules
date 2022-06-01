package errors

import "errors"

var (
	// ErrNotFound is returned when a resource is not found.
	ErrNotFound = errors.New("not found")
	// ErrTimeout is returned when a timeout occurs.
	ErrTimeout = errors.New("timeout")
)
