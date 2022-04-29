package main

import "fmt"

func main() {
	blacklist := func(name string) bool {
		return name == "Anjing"
	}
	registerUser("eko", blacklist)
	
	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})
}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

