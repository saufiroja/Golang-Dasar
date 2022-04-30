package main

import "fmt"

func main() {
	person := Person {
		Name: "Oja",
	}
	SayHai(person)

	animal := Animal {
		Name: "Kucing",
	}

	SayHai(animal)
}

// test 1
type Person struct {
	Name string
}
func (person Person) GetName() string {
	return person.Name
}

// test 2
type Animal struct {
	Name string
}
func (animal Animal) GetName() string {
	return animal.Name
}

// interface
type HasName interface {
	GetName() string
}

func SayHai(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}