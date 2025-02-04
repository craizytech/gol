package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

// define a struct
type Person struct {
	Name        string
	Age         int
	Nationality []string
}

// define a method with a receiver of type Person
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func ex13() {
	p := Person{Name: "Eammon", Age: 22, Nationality: []string{"Kenyan"}}
	p1 := Person{Name: "Eammon", Age: 22, Nationality: []string{"Kenyan"}}

	p2 := Person{Name: "Gibson", Age: 24, Nationality: []string{"Kenyan"}}

	fmt.Println(cmp.Equal(p, p1))
	fmt.Println(cmp.Equal(p, p2))
}
