package main

import (
	"fmt"
	"log"
	"strconv"
)

// Declare a Book type which satisfies the Stringer Interface
type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

// Declare a count type that satisfies the fmt.Stringer Interface

type Count int

func (c Count) String() string {
	return strconv.Itoa(int(c))
}

// Declare a WriteLog() function which only takes any object that
// satisfies the fmt.Stringer interface as a parameter
func WriteLog(s fmt.Stringer) {
	log.Print(s.String())
}

func ex5() {
	// initialize a count object and pass it to WriteLog()
	book := Book{"Alice in WonderLand", "Lewis Carrol"}
	WriteLog(book)

	//initialize a count object and pass it to the WriteLog()
	count := Count(3)

	WriteLog(count)
}
