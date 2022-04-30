package main

import (
	"fmt"
)

func main() {
	book := "The colour of magic"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
	//uint8 is a byte

	// strings in go are immutable, meaning you can not change them once created
	// book[0] = 116

	//Slice (start, end), 0 based, 1/2 empty range
	fmt.Println(book[4:11])

	// Slice (no end)
	fmt.Println(book[4:])

	// slice (no start)
	fmt.Println(book[:4])

	//Use + to concatenate strings together
	fmt.Println("t" + book[1:])

	//Unicode
	fmt.Println("It was ½ price!")

	// Multiline

	poem := `
	The road goes even on
	Down from the door where it began
	...
	`

	fmt.Println(poem)

}
