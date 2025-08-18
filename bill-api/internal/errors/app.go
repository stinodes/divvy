package internalerrors

import "errors"

var (
	ErrBadInput     = errors.New("bad input")
	ErrNotFound     = errors.New("not found")
	ErrForbidden    = errors.New("forbidden")
	ErrUnauthorized = errors.New("unauthorized")
)
