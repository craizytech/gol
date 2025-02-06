package main

import (
	"fmt"
	"sort"
)

type people struct {
	name  string
	email string
	dob   dob
}
type dob struct {
	day   int
	month int
	year  int
}

func ex17() {
	members := make(map[int]people)

	members[1] = people{
		name:  "Mary Smith",
		email: "marysmith@example.com",
		dob: dob{
			day:   17,
			month: 3,
			year:  1990,
		},
	}
	members[2] = people{
		name:  "John Smith",
		email: "johnsmith@example.com",
		dob: dob{
			day:   9,
			month: 12,
			year:  1988,
		},
	}

	members[3] = people{
		name:  "Janet Doe",
		email: "janetdoe@example.com",
		dob: dob{
			day:   1,
			month: 12,
			year:  1988,
		},
	}

	members[4] = people{
		name:  "Adam Jones",
		email: "adamjones@example.com",
		dob: dob{
			day:   19,
			month: 8,
			year:  2001,
		},
	}

	for k, v := range members {
		fmt.Println(k, v.name, v.email, v.dob)
	}

	fmt.Println("-------------")

	// Sorting the map of structs
	// Printing the contents of the map in a certain order

	// get all the keys in members
	var keys []int
	for k := range members {
		keys = append(keys, k)
	}

	// sort the keys in ascending order
	sort.Ints(keys)

	//copy the value in members to the slice
	var sliceMembers []people
	for _, k := range keys {
		sliceMembers = append(sliceMembers, members[k])
	}

	// fmt.Println(sliceMembers)

	fmt.Println("-----------------")

	sort.SliceStable(sliceMembers, func(i, j int) bool {
		//compare year
		if sliceMembers[i].dob.year != sliceMembers[j].dob.year {
			return sliceMembers[i].dob.year < sliceMembers[j].dob.year
		}

		// compare month
		if sliceMembers[i].dob.month != sliceMembers[j].dob.month {
			return sliceMembers[i].dob.month < sliceMembers[j].dob.month
		}

		//compare day
		return sliceMembers[i].dob.day < sliceMembers[j].dob.day
	})

	for _, v := range sliceMembers {
		fmt.Println(v.name, v.dob)
	}

}
