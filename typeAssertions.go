package main

import "fmt"

func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int) // panic
	// fmt.Println(resultInt)

	// switch type assertions
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}
}

func random() interface{} {
	return "Ok"
}