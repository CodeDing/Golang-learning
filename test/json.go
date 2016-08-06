package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name    string
	Age     int
	Guake   bool
	Classes []string
	Price   float32
}

func main() {
	st := Student{
		"Xiao Ming",
		16,
		true,
		[]string{"Math", "English"},
		9.99,
	}
	st1 := Student{
		"Xiao Ming",
		16,
		true,
		[]string{"Math", "English"},
		9.99,
	}
	str := make([]Student, 0)
	str = append(str, st)
	str = append(str, st1)
	b, err := json.Marshal(st)
	c, err := json.Marshal(str)
	if err != nil {
		fmt.Println("encoding failed")
	} else {
		fmt.Println("encoding data: ")
		fmt.Println(string(b))
	}
	fmt.Println(string(c))
}
