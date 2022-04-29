package main

import "fmt"

func main() {
	var name [3]string
	name[0] = "Muhammad"
	name[1] = "Saufi"
	name[2] = "Roja"

	fmt.Println(len(name[0]))
	fmt.Println(name[1])
	fmt.Println(name[2])

	var value = [3]int{
		90,
		80,
		70,
	}
	value[0] = 100
	
	fmt.Println(value)
	fmt.Println(len(value))
}