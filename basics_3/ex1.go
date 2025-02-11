package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var balance int

var mutex = &sync.Mutex{}

func credit() {
	for i := 0; i < 5; i++ {
		mutex.Lock()
		balance += 100
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Println("After Crediting, balance is", balance)
		mutex.Unlock()
	}

}

func debit() {
	for i := 0; i < 5; i++ {
		mutex.Lock()
		balance -= 100
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Println("After Debiting, balance is", balance)
		mutex.Unlock()
	}
}

func ex1() {
	balance = 200
	fmt.Println("Initial balance is", balance)
	go credit()
	go debit()
	fmt.Scanln()
}
