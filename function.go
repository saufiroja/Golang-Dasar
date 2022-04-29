package main

import "fmt"


func main() {
	// funtion with parameter
	for i := 0; i < 5; i++{
		sayHello("Muhammad", "Saufi")
	}

	// return function
	result := getHello("Muhammad Saufi Roja")
	fmt.Println(result)

	// multiple function
	firstName, _ ,lastName := getFullName()
	fmt.Println(firstName, lastName)

	// named return value
	fmt.Println(getCompleteName())

	// variadic function
	total := sumAll(5, 5, 5, 5, 5)
	fmt.Println(total)

	numbers := []int{5, 5, 5, 5, 10}
	total = sumAll(numbers...)
	fmt.Println(total)


	// function value
	sayGoodBye := getGoodBye
	value := sayGoodBye("Eko")
	fmt.Println(value)
}


// funtion with parameter
func sayHello(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

// return function
func getHello(name string) string {
	if name == ""{
		return "Hello bro"
	} else {
		return "Hello " + name
	}
}

// multiple function
func getFullName() (string, string, string){
	return "Eko", "Kurniawan", "Khannedy"
}

// named return value
func getCompleteName () (value1, value2, value3 string) {
	value1 = "Muhammad"
	value2 = "Saufi"
	value3 = "Roja"
	return
}

// variadic function
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

// function value
func getGoodBye(name string) string {
	return "GoodBye " + name
}