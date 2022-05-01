package main

import (
	"fmt"
	"sort"
)

func main() {
	user := []User{
		{"Eko", 30},
		{"Oja", 20},
		{"Budi", 26},
	}

	sort.Sort(UserSlice(user))
	fmt.Println(user)
}

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}