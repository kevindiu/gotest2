package testdata

func Max[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type Container[T any] struct {
	Value T
}

func (c *Container[T]) Get() T {
	return c.Value
}

func Map[T any, R any](s []T, f func(T) R) []R {
	r := make([]R, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}
