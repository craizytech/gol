package main

import "fmt"

func filter(arr []int, cond func(int) bool) []int {
	result := []int{}

	for _, v := range arr {
		if cond(v) {
			result = append(result, v)
		}
	}
	return result

}

func ex8() {
	a := []int{1, 2, 3, 4, 5, 6, 7}

	evens := filter(a, func(num int) bool {
		return num%2 == 0
	})

	fmt.Println(evens)
}
