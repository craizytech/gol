package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
	name   string
}

type Square struct {
	length float64
	name   string
}

type Triangle struct {
	base   float64
	height float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) Circumfrence() float64 {
	return 2 * math.Pi * c.radius
}

func (s Square) Area() float64 {
	return math.Pow(s.length, 2)
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func calculateArea(listOfShapes []Shape) {
	for _, shape := range listOfShapes {
		fmt.Println("Area of Shape is ", shape.Area())
	}
}

func ex7() {
	c1 := Circle{radius: 5, name: "c1"}
	s1 := Square{length: 6, name: "s1"}
	t1 := Triangle{base: 4, height: 5}

	shapes := []Shape{c1, s1, t1}
	calculateArea(shapes)
}
