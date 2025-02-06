package main

import (
	"errors"
	"fmt"
	"strings"
)

// function to perform addition
// func add(a int, b int) int {
// 	return a + b
// }

// shorthand syntax for parameters of the same type
func perimeter(length, width int) int {
	return 2 * (length + width)
}

// function with multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Division by Zero")
	} else {
		return a / b, nil
	}
}

// function with named return values
func rectangleArea(length, width float64) (area float64) {
	area = length * width
	return
}

// variadic functions
func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

// functions as a value
func multiply(a, b int) int {
	return a * b
}

// functions as a parameter
func operate(a, b int, op func(int, int) int) int {
	return op(a, b)
}
func add(x, y int) int {
	return x + y
}

// Closure Functions
func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

// High Order Functions - functions that returns other functions
func multiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

// Pointers and Functions
func increment(n *int) {
	*n++
}

// function to do sth
func doSth() (int, bool) {
	return 5, false
}

func ex1() {
	// defer fmt.Println("After executing everything") // defered function
	// fmt.Println(perimeter(4, 5))
	// fmt.Println(divide(6, 0))
	// fmt.Println(rectangleArea(4, 4))
	// fmt.Println(sum(3, 4, 5, 6, 7))

	// // functions as first class citizens
	// var op func(int, int) int
	// op = multiply
	// fmt.Println(op(3, 4))

	// // functions as arguments/paramenters
	// fmt.Println(operate(5, 3, add))

	// // lambda function
	// res := func(a, b int) int {
	// 	return a + b
	// }(3, 4)

	// fmt.Println(res)

	// // closure function
	// next := counter()
	// fmt.Printf("%T\n", next)
	// fmt.Println(next())
	// fmt.Println(next())

	// // higher order functions
	// mul := multiplier(9)
	// fmt.Println(mul(5))

	// x := 10
	// increment(&x)
	// fmt.Println(x)

	// // keeping variable scopes tight
	// if v, err := doSth(); !err {
	// 	fmt.Println("An error Occured")
	// } else {
	// 	fmt.Printf("we have a number %d", v)
	// }

	// num := 5
	// dayOfWeek := ""

	// switch num {
	// case 1:
	// 	dayOfWeek = "Monday"
	// case 2:
	// 	dayOfWeek = "Tuesday"
	// case 3:
	// 	dayOfWeek = "Wednesday"
	// case 4:
	// 	dayOfWeek = "Thursday"
	// case 5:
	// 	dayOfWeek = "Friday"
	// case 6:
	// 	dayOfWeek = "Saturday"
	// case 7:
	// 	dayOfWeek = "Sunday"
	// default:
	// 	dayOfWeek = "--error--"
	// }

	// fmt.Println(dayOfWeek)

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// max := 100
	// a, b := 0, 1

	// for b <= max {
	// 	fmt.Println(b)
	// 	a, b = b, a+b
	// }

	for {
		fmt.Println("enter QUIT to exit")
		var input string
		fmt.Print("Please enter a string: ")
		fmt.Scanf("%s", &input)

		if strings.ToUpper(input) == "QUIT" {
			break
		}
	}
}
