package main

import (
	"fmt"
	"strings"
)

func main() {
	policy := "field1:tcp,field2:16888,field3:10.10.24.200,field4:16888"
	policy_strs := strings.Split(policy, ",")
	fmt.Println(len(policy_strs))
	fmt.Println(policy_strs[0])
	fmt.Println(policy_strs[1])
	fmt.Println(policy_strs[2])
	fmt.Println(policy_strs[3])
	for _, policy_str := range policy_strs {
		fmt.Println(policy_str)
		fmt.Println(strings.Split(policy_str, ":")[1])
	}
}
