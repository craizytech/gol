package main

import (
	"errors"
)

func insert(orig []int, value int, index int) ([]int, error) {
	if index < 0 {
		return nil, errors.New("The index cannot be less than 0")
	}

	if index >= len(orig) {
		return append(orig, value), nil
	}

	orig = append(orig[:index+1], orig[index:]...)
	orig[index] = value

	return orig, nil
}

func delete(orig []int, index int) ([]int, error) {
	if index < 0 || index >= len(orig) {
		return nil, errors.New("Index is out of range")
	}

	orig = append(orig[:index], orig[index+1:]...)

	return orig, nil
}
