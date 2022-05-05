package main

import (
	"fmt"
	"math"
)

// Shape is an interface
type Shape interface {
	Area() float64
}

// Square is a square
type Square struct {
	Length float64
}

// Area returns the area of the Square
func (s *Square) Area() float64 {
	return s.Length * s.Length
}

// Circle is a circle
type Circle struct {
	Radius float64
}

// Area returns the circle of the Square
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//sumAreas returns the sum of alll areas in the slices
func sumAreas(shapes []Shape) float64 {
	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func main() {
	s := &Square{20}
	fmt.Println(s.Area())

	c := &Circle{10}
	fmt.Println(c.Area())

	shapes := []Shape{s, c}
	sa := sumAreas(shapes)
	fmt.Println(sa)


}

