package main

import "fmt"

func main(){
	const firstName = "Muhammad Saufi"
	const lastName = "Roja"

	fmt.Println(firstName, lastName)

	// multiple constant
	const (
		country = "Indonesia"
		age = 77
	)
	fmt.Println(country)
	fmt.Println(age)
}