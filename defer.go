package main

import "fmt"

func main() {
	runApp(10)
}

func logging() {
	fmt.Println("Selesai manggil function")
}

func runApp(value int) {
	defer logging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println(result)
}