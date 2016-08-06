package main

import "fmt"

func main() {
	x := [3]int{1, 2, 3}

	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr) //prints [7 2 3]
	}(x)
	fmt.Println(x) //prints [1 2 3] (not ok if you need [7 2 3])

	var n [10]int
	var i, j int

	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
	for _, v := range n {
		fmt.Printf("v = %d\n", v)
	}

}
