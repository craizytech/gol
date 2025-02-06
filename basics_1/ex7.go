package main

import "fmt"

func doSomething() int {
	return 5
}

func ex7() {
	var i func() int

	i = doSomething

	fmt.Println(i())
}
