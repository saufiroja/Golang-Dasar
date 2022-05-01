package main

import "fmt"

func main() {
	eko := Man{"Eko"}
	eko.Married()

	fmt.Println(eko.Name)
}

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}