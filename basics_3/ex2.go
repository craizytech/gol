package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var balance1 int64

func credit1() {
	for i := 0; i < 10; i++ {
		// Adds 100 to the balance atomically
		atomic.AddInt64(&balance1, 100)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Microsecond)
	}
}

func debit1() {
	for i := 0; i < 10; i++ {
		// deducts 100 from the balance atomically
		atomic.AddInt64(&balance1, -100)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Microsecond)
	}
}

func ex2() {
	balance1 = 200
	fmt.Println("Initial balance is: ", balance1)
	go credit1()
	go debit1()
	fmt.Scanln()
	fmt.Println(balance1)
}
