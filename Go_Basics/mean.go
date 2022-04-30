package main

import (
	"fmt"
)

func main() {
	// go's type system is very strict, it does not allow to add ints to a float.

	// := operator presumes the type of the object, for example below go will know that we
	// are dealing with floats
	x, y := 1.0, 2.0

	// %v verb will print the go object, the %T verb will print the type
	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	mean := (x + y) / 2.0
	fmt.Printf("result: %v, type of %T\n", mean, mean)
}
