package main

import "fmt"

type laptop struct {
	cpu          string
	ram          int
	storage      int
	manufacturer string
}

func (l *laptop) upgradeStorage(size int) {
	l.storage += size
}

func ex12() {
	thinkpad := laptop{"i5-1240p", 16, 256, "lenovo"}
	fmt.Println(thinkpad.storage)
	thinkpad.upgradeStorage(100)
	fmt.Println(thinkpad.storage)
}
