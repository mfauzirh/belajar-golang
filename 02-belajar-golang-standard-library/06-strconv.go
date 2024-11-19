package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("True")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}

	result2, err2 := strconv.Atoi("100")
	if err2 != nil {
		fmt.Println(err2.Error())
	} else {
		fmt.Println(result2)
	}
}
