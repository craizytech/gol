package main

import (
	"encoding/json"
	"fmt"
)

type Details struct {
	Height int
	Weight float32
}
type People struct {
	Firstname string
	Lastname  string
	Details   Details
}

func ex19() {
	var persons []People
	jsonString :=
		`
	[
	{
	"firstname":"Eammon",
	"last name":"Kiprotich",
	"details": {
	"height":175,
	"weight":70.0}
	},
	{
	"firstname":"Mickey",
	"lastname":"Mouse",
	"details": {
	"height":105,
	"weight":85.5}
	}
	]
	`
	json.Unmarshal([]byte(jsonString), &persons)

	for _, person := range persons {
		fmt.Println(person.Firstname)
		fmt.Println(person.Lastname)
		fmt.Println(person.Details.Height)
		fmt.Println(person.Details.Weight)
	}

}
