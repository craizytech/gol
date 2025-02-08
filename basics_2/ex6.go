package main

import "fmt"

type DigitsCounter interface {
	CountOddEven() (int, int)
}

type DigitString string

func (ds DigitString) CountOddEven() (int, int) {
	odds, evens := 0, 0

	for _, digit := range ds {
		if digit%2 == 0 {
			evens++
		} else {
			odds++
		}
	}
	return odds, evens
}

func ex6() {
	var d DigitsCounter
	s := DigitString("0123456789")
	d = s
	fmt.Println(d.CountOddEven())
}
