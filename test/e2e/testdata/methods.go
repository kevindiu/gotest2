package testdata

type Calculator struct {
	base int
}

func (c *Calculator) Add(n int) int {
	return c.base + n
}

type Greeter struct {
	Name string
}

func (g Greeter) Greet() string {
	return "Hello " + g.Name
}
