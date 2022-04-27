package main

import "fmt"

func main(){
	// 1
	var name string
	name = "Muhammad Saufi Roja"
	fmt.Println(name)

	// 2
	var name1 = "Muhammad Saufi Roja"
	fmt.Println(name1)

	// 3
	name2 := "Muhammad Saufi Roja"
	name2 = "Indonesia"
	fmt.Println(name2)

	// Multiple variable
	var (
		firstName = "Muhammad Saufi"
		lastName = "Roja"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}