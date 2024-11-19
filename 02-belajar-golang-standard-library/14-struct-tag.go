package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `required:"true" max:"10"` // this is struct tag
	Age  int    `required:"true" min:"0"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)

	fmt.Println("Type name", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Println(valueField.Name, " - ", valueField.Type)
		fmt.Println(valueField.Tag.Get("required"))
		fmt.Println(valueField.Tag.Get("max"))
		fmt.Println(valueField.Tag.Get("min"))
	}
}

func isValid(value any) (result bool) {
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {
	person := Person{"John", 12}

	readField(person)
	fmt.Println("\n\n")
	fmt.Println(isValid(person))
}
