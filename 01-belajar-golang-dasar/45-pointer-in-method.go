package main

import "fmt"

type Man struct {
	Name string
}

func (m *Man) Married() {
	m.Name = "Mr. " + m.Name
}

func main() {
	eko := Man{"Eko"}
	eko.Married()

	fmt.Println(eko.Name)
}
