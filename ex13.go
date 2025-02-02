package main

import "fmt"

// define a struct
type Person struct {
	Name string
	Age  int
}

// define a method with a receiver of type Person
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func ex13() {
	p := Person{Name: "Eammon", Age: 22}
	p.Greet() // calling the method
}
