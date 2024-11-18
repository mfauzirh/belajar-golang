package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi adalah 0")
	}
	return nilai / pembagi, nil
}

func main() {
	hasil, err := Pembagian(20, 0)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(hasil)
	}
}
