package main

import "errors"

func MultiError(n int) (int, error, error) {
	if n < 0 {
		return 0, errors.New("negative"), nil
	}
	if n > 100 {
		return 0, nil, errors.New("too large")
	}
	return n, nil, nil
}
