package main

import (
	"fmt"
	"math"
)

type point struct {
	x float32
	y float32
	z float32
}

func newPoint(x, y, z float32) *point {
	p := point{x: x, y: y, z: z}
	return &p
}

// length is a method with a value receiver
func (p point) length() float64 {
	return math.Sqrt(
		(math.Pow(float64(p.x), 2) +
			math.Pow(float64(p.y), 2) +
			math.Pow(float64(p.z), 2)))
}

// move is a method with a pointer receiver
func (p *point) move(deltax, deltay, deltaz float32) {
	p.x += deltax
	p.y += deltay
	p.z += deltaz
}

func ex10() {

	// pt4 := newPoint(3.2, 1.5, 6.2)
	// pt5 := pt4

	// fmt.Println("pt4: ", pt4)
	// fmt.Println("pt5: ", pt5)

	// fmt.Println("-----------")
	// pt5.y = 3.14

	// fmt.Println("pt4: ", pt4)
	// fmt.Println("pt5: ", pt5)

	// fmt.Println("-----------")
	// pt6 := *pt4
	// fmt.Println(pt6)
	// pt6.y = 9.8
	// fmt.Println(pt6)

	pt7 := newPoint(3, 4, 12)
	fmt.Println(pt7.length())

	fmt.Println(*pt7)

	pt7.move(1, 2, 3)
	fmt.Println(*pt7)
}
