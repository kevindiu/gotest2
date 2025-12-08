package main

import "fmt"

// Container is a generic interface
type Container[T any] interface {
	Get() T
	Put(T)
}

// Box implements Container for any type
type Box[T any] struct {
	val T
}

func (b *Box[T]) Get() T {
	return b.val
}

func (b *Box[T]) Put(v T) {
	b.val = v
}

// ProcessContainer is a function that takes a generic interface
func ProcessContainer[T any](c Container[T], val T) {
	c.Put(val)
	fmt.Println(c.Get())
}
