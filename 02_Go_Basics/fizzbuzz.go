package main

import (
	"fmt"
)

func main() {
	a := "fizz"
	b := "buzz"

	for i := 1; i < 21; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(a, b)
		} else if i%3 == 0 {
			fmt.Println(a)
		} else if i%5 == 0 {
			fmt.Println(b)
		} else {
			fmt.Println(i)
		}
	}
}
