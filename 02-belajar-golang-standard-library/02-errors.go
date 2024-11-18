package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "Eko" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := GetById("John")

	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation Error Occur")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not Found Error Occur")
		} else {
			fmt.Println("Unknown error")
		}
	}
}
