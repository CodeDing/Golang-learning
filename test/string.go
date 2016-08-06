package main

import (
	"fmt"
	"strings"
)

const (
	PRIVATE_IP = "10."
)

func main() {
	str1 := "10.2.3.4"
	str2 := "192.34.32.23"
	if strings.HasPrefix(str1, PRIVATE_IP) {
		fmt.Println("HasPrefix")
	}
	if strings.HasPrefix(str2, PRIVATE_IP) {
		fmt.Println("HasPrefix")
	}
	ip := make([]string, 0)
	ip = append(ip, "12.23.34.56")
	fmt.Println(strings.Join(ip, "','"))
	fmt.Println(ip)
	strs := make([]string, 0)
	strs = append(strs, "")
	fmt.Println(str_len(strs))

}

func str_len(strs []string) int {
	if len(strs) == 0 {
		fmt.Println("strs is empty")
	}
	return len(strs)
}
