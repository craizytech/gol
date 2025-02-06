package main

import "fmt"

type MyInt int

func (m MyInt) Double() MyInt {
	return m * 2
}

func ex14() {
	num := MyInt(10)
	fmt.Println(num.Double())
}
