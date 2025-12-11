package testdata

func Sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func PrintAll(prefix string, args ...interface{}) {
	// do nothing
}
