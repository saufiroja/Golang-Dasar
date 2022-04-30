package main

import (
	"errors"
	"fmt"
)

func main() {
	value, err := Pembagian(6,0)
	if err == nil {
		fmt.Println("Hasil", value)
	} else {
		fmt.Println("Error", err)
	}
}

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan nol")
	} else {
		return nilai / pembagi, nil
	}
}