package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var balance2 int64
var mutex2 = &sync.Mutex{}

func credit2(wg *sync.WaitGroup) {
	// notify the WaitGroup when we are done
	defer wg.Done()

	for i := 0; i < 10; i++ {
		// adds 100 to balance atomically
		mutex2.Lock()
		atomic.AddInt64(&balance2, 100)
		fmt.Println("Credit: Balance after adding 100 =", balance2)
		mutex2.Unlock()

		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func debit2(wg *sync.WaitGroup) {
	// notify the WaitGroup when we are done
	defer wg.Done()

	for i := 0; i < 5; i++ {
		// deducts -100 from balance atomically
		mutex2.Lock()
		atomic.AddInt64(&balance2, -100)
		fmt.Println("Debit: Balance after subtracting 100 =", balance2)
		mutex2.Unlock()

		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func ex3() {
	var wg sync.WaitGroup

	balance2 = 200
	fmt.Println("Initial balance is", balance2)

	wg.Add(1) // add 1 to the WaitGroup counter
	go credit2(&wg)

	wg.Add(1) // add 1 to the WaitGroup counter
	go debit2(&wg)

	wg.Wait() // blocks until WaitGroup counter is 0
	fmt.Println("Final balance is", balance2)
}
