package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	// for dengan statement
	for i := 1; i < 10; i++ {
		fmt.Println("Perulangan ke ", i)
	}

	slice := []string{"Muhammad", "Saufi", "Roja"}
	for i := 0; i < len(slice); i++{
		fmt.Println(slice[i])
	}

	// for range
	names := []string{"Eko", "Kurniawan", "Khannedy"}
	for i , value := range names { 					//index bisa di ganti _ apabila tidak digunakan index nya
		fmt.Println("index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Oja"
	person["title"] = "backend"
	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}