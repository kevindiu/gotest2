package main

func GenericSum[T int | float64](a, b T) T {
	return a + b
}

type MyList[T any] struct {
	items []T
}

func (l *MyList[T]) Add(item T) {
	l.items = append(l.items, item)
}

func (l *MyList[T]) Get(index int) T {
	return l.items[index]
}

// Swap returns the two arguments swapped.
func Swap[T any](a, b T) (T, T) {
	return b, a
}
