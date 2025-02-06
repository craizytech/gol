package main

import (
	"fmt"
	"sort"
)

func ex15() {
	heights := map[string]int{
		"peter": 150,
		"joan":  168,
		"janet": 175,
	}

	// fmt.Printf("The len of the map is: %d\n", len(heights))

	// if val, exists := heights["Peter"]; exists {
	// 	fmt.Println(val)

	// } else {
	// 	fmt.Println("The key does not exist")
	// }
	var keys []string

	for key := range heights {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, v := range keys {
		fmt.Println(v, heights[v])
	}
}
