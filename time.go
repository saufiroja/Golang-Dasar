package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2001, time.May, 2, 5, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local())

	format := "2006-01-02"
	parse, _ := time.Parse(format, "2020-02-05")
	fmt.Println(parse)
}