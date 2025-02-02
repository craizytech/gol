package main

import "fmt"

func ex3() {
	for pos, char := range "こんにちは世界" {
		fmt.Printf("%c at byte location %d\n", char, pos)
	}
}
