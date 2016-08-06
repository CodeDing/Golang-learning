package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	printSlice(numbers)

	fmt.Println("numbers == ", numbers)
	fmt.Println("numbers[:3] ", numbers[:3])

	numbers1 := make([]int, 0, 5)
	numbers1 = append(numbers1, 1, 2)
	printSlice(numbers1)
	numbers1 = append(numbers1, 1, 2)
	numbers1 = append(numbers1, 1, 2)
	printSlice(numbers1)

	numbers3 := numbers[2:5]
	printSlice(numbers3)

	var numbers2 []int
	printSlice(numbers2)
	numbers2 = append(numbers2, 0)
	printSlice(numbers2)

	a := [10]int{1, 2, 3, 4, 5, 7, 8, 9, 0}
	fmt.Println(len(a))
	sa := a[2:7]
	fmt.Println(sa)
	sa = append(sa, 100)
	sa = append(sa, 1)
	fmt.Println(sa)
	sb := sa[3:8]
	sb[0] = 99
	fmt.Println(a)
	fmt.Println(sa)
	fmt.Println(sb)
}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n ", len(x), cap(x), x)
}
