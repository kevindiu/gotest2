package testdata

import "errors"

func DivMod(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("divide by zero")
	}
	return a / b, a % b, nil
}

func NamedReturn(x int) (res int, err error) {
	if x < 0 {
		err = errors.New("negative")
		return
	}
	res = x * 2
	return
}
