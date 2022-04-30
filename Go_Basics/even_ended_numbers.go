package main

import (
	"fmt"
)

func main() {

	count := 0

	// for every pair of 4 digit numbers
	for i := 1000; i <= 9999 ; i++{
		for j := i;  j<=9999; j++ { // don't count twice
			n:= i*j

			//if a * b even ended
			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s) -1]{
			//count - count +1
			count++
		}

	
	}
}
	
	fmt.Println(count)
}