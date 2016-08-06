package main

import "fmt"

func main() {
	var x interface{}
	x = "23"
	switch i := x.(type) {
	case nil:
		fmt.Printf(" x的类型：%T", i)
	case int:
		fmt.Printf("x是int类型")
	case float64:
		fmt.Printf("x是float64类型")
	case func(int) float64:
		fmt.Printf("x是func(int)类型")
	case bool, string:
		fmt.Printf("x是bool或是string类型")
	default:
		fmt.Printf("未知型")
	}
}
