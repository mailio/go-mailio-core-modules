package errors

import "errors"

var (
	// ErrNotFound is returned when a resource is not found.
	ErrNotFound = errors.New("not found")
	// ErrTimeout is returned when a timeout occurs.
	ErrTimeout = errors.New("timeout")
	// ErrInvalidFormat is returned when a resource is not in the expected format.
	ErrInvalidFormat = errors.New("invalid format")
	// operation not authorized
	ErrUnauthorized = errors.New("unauthorized")
	// ErrNotAuthorized - authorization failed
	ErrNotAuthorized = errors.New("not authorized")
	// ErrUserExists - user already exists
	ErrUserExists = errors.New("user exists")
	// ErrUnimoplemented - unimplemented
	ErrUnimplemented = errors.New("unimplemented")
	// ErrBadRequest - bad request (or invalid input data)
	ErrBadRequest = errors.New("bad request")
	// ErrQuotaExceeded - quota exceeded
	ErrQuotaExceeded = errors.New("quota exceeded")
	// expired trial
	ErrTrialExpired = errors.New("trial has expired")
	// ErrRateLimitExceeded - rate limit exceeded
	ErrRateLimitExceeded = errors.New("rate limit exceeded")
)
