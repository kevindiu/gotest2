package main

// VariadicSum sums an arbitrary number of integers.
func VariadicSum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

// ChanSquare reads from input channel, squares numbers, and writes to output channel.
// Returns the output channel.
func ChanSquare(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

// MapKeys returns the keys of a map.
func MapKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// FuncFactory returns a function that adds n to its input.
func FuncFactory(n int) func(int) int {
	return func(x int) int {
		return x + n
	}
}

// DeferredExecution simulates a function using defer (simple check).
func DeferredExecution() (msg string) {
	defer func() {
		msg += " world"
	}()
	return "hello"
}
