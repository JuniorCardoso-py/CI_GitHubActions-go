package main

import "fmt"

func main() {
	a, b := 25, 50
	total := Sum(a, b)

	fmt.Printf("O resultado da soma entre %v e %v = %v", a, b, total)
}

func Sum(a, b int) int {
	return a + b
}
