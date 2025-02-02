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
	// createStruct := point{x: 5, y: 6, z: 7}

	// pt1 := point{x: 2.2, z: 3.3}
	// fmt.Println(pt1)

	// new1 := createStruct
	// new1 = point{x: 0}

	// fmt.Println(createStruct, new1)

	pt4 := newPoint(3.2, 1.5, 6.2)
	pt5 := pt4

	fmt.Println("pt4: ", pt4)
	fmt.Println("pt5: ", pt5)

	fmt.Println("-----------")
	pt5.y = 3.14

	fmt.Println("pt4: ", pt4)
	fmt.Println("pt5: ", pt5)

	fmt.Println("-----------")
	pt6 := *pt4
	fmt.Println(pt6)
	pt6.y = 9.8
	fmt.Println(pt6)

}
