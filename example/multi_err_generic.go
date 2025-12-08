package main

import "errors"

// GenericMultiError is a generic function that returns a value and two errors.
func GenericMultiError[T any](val T, n int) (T, error, error) {
	if n < 0 {
		return val, errors.New("negative"), nil
	}
	if n > 100 {
		var zero T
		return zero, nil, errors.New("too large")
	}
	return val, nil, nil
}
