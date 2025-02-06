package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func ex6() {
	x := 5
	y := 6
	swap(&x, &y)
	fmt.Println(x, y)
}
