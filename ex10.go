package main

import "fmt"

type point struct {
	x float32
	y float32
	z float32
}

func newPoint(x, y, z float32) *point {
	p := point{x: x, y: y, z: z}
	return &p
}

func ex10() {
	createStruct := point{x: 5, y: 6, z: 7}

	new1 := createStruct
	new1 = point{x: 0}

	fmt.Println(createStruct, new1)

	// pt4 := newPoint(3.2, 1.5, 6.2)

	// fmt.Println(pt4.x)
}
