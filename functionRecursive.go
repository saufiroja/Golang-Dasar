package main

import "fmt"

func main() {
	// function factorial for loop
	loop := factorialLoop(5)
	fmt.Println(loop)

	// function recursive factorial loop
	recurive := factorialRecursive(5)
	fmt.Println(recurive)
}

// function factorial for loop
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

// function recursive factorial loop
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value - 1)
	}
}