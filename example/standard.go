package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) Greet(msg string) string {
	return fmt.Sprintf("%s, I am %s", msg, p.Name)
}

// DivMod returns quotient and remainder
func DivMod(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, fmt.Errorf("division by zero")
	}
	return a / b, a % b, nil
}

func Add(a, b int) int {
	return a + b
}
