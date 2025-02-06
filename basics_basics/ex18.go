package main

import (
	"encoding/json"
	"fmt"
)

type Peoples struct {
	Firstname string
	Lastname  string
}

func ex18() {
	var person People

	jsonString := `{
	"firstname":"Eammon",
	"lastname":"Kiprotich"
	}`

	byteSlice := []byte(jsonString)
	println(byteSlice)

	err := json.Unmarshal([]byte(jsonString), &person)

	if err == nil {
		fmt.Println(person.Firstname)
		fmt.Println(person.Lastname)
	} else {
		fmt.Println(err)
	}
}
