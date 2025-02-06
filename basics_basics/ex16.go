package main

import (
	"fmt"
	"sort"
)

// define a struct to hold key-value pairs
type kv struct {
	key   string
	value int
}

// Define a type alias for a slice of kv structs
type kvPairs []kv

// Implement sort.Interface for kvPairs
func (p kvPairs) Len() int {
	// returns the lenght of the collection
	return len(p)
}

func (p kvPairs) Less(i, j int) bool {
	// indicates the first value (height) must be smaller
	// than the second value
	return p[i].value < p[j].value
}

func (p kvPairs) Swap(i, j int) {
	// swaps the items in the collection
	p[i], p[j] = p[j], p[i]
}

func ex16() {
	heights := map[string]int{
		"Peter": 170,
		"Joan":  168,
		"Janet": 175,
	}

	// convert the map to a slice of kv structs
	p := make(kvPairs, 0, len(heights))

	for k, v := range heights {
		p = append(p, kv{k, v})
	}

	// sort the slice based on values
	sort.Sort(p)

	// Print sorted results
	for _, v := range p {
		fmt.Println(v.key, v.value)
	}

}
