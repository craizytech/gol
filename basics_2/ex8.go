package main

import (
	"fmt"
	"time"
)

func say(s string, times int) {
	for i := 0; i < times; i++ {
		// inject a 100ms delay
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(i, s)
	}
}

func ex8() {
	go say("Hello", 3)
	go say("World", 2)
	fmt.Scanln()
}
