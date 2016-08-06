package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "01__02-2006 3.04.05 PM"

	fmt.Println(time.Now().Format(layout))
}
