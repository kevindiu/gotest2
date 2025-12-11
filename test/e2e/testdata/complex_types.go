package testdata

type ComplexData struct {
	ID   int
	Tags []string
	Meta map[string]interface{}
}

func ProcessData(d ComplexData) (result map[string]int) {
	return map[string]int{"count": len(d.Tags)}
}

func AsyncProcess(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * 2
	}
}

func PointerToPointer(p **int) *int {
	if p == nil {
		return nil
	}
	return *p
}
