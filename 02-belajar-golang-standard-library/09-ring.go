package main

import (
	"container/ring"
	"fmt"
)

func main() {
	var data *ring.Ring = ring.New(5)

	data.Value = "Value 1"

	data = data.Next()
	data.Value = "Value 2"

	data = data.Next()
	data.Value = "Value 3"

	data = data.Next()
	data.Value = "Value 4"

	data = data.Next()
	data.Value = "Value 5"

	data.Do(func(value any) {
		fmt.Println(value)
	})
}