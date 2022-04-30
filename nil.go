package main

import "fmt"

func main() {
	var data = NewMap("oja")
	if data == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(data)
	}
}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}