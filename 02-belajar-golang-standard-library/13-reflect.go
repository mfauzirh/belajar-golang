package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string
}

type Person struct {
	Name string
	Age  int
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)

	fmt.Println("Type name", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Println(valueField.Name, " - ", valueField.Type)
	}
}

func main() {
	readField(Sample{"Joko"})
	readField(Person{"Budi", 20})
}
