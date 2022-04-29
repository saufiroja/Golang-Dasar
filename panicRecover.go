package main

import "fmt"

func main() {
	runApplication(false)
}

func endApp() {
	message := recover()
	if message != nil {		
		fmt.Println("Error dengan", message)
	}
	fmt.Println("End App")
}

func runApplication(error bool) {
	defer endApp()
	if error {
		panic("Error")
	}
	fmt.Println("App berjalan")
}