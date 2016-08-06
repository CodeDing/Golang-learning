package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	str := "Dog,Apple,Cat,Boy"

	strSlice := strings.Split(str, ",")

	fmt.Println("Before : ", strSlice)

	sortableStrSlice := sort.StringSlice(strSlice)

	sort.Sort(sortableStrSlice)

	fmt.Println("After : ", sortableStrSlice)
}
