package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (p Person) GetName() string {
	return p.Name
}

func main() {
	person := Person{"Eko"}
	SayHello(person)
}
