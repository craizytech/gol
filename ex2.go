package main

import "fmt"

func ex2() {
	var OS [3]string
	OS[0] = "IOS"
	OS[1] = "Android"
	OS[2] = "Windows"

	for _, v := range OS {
		fmt.Println(v)
	}
}
