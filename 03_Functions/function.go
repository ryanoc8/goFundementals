package main

import (
	"fmt"
)

//add adds a to b
func add(a, b int) int {
	return a + b
}

func divmod(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	val := add(1, 2)
	fmt.Println(val)

	div, mod := divmod(7, 2)
	fmt.Printf("div=%d: mod%d\n", div, mod)
}
