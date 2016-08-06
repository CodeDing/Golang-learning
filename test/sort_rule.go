package main

import (
	"fmt"
	"sort"
)

type Rule struct {
	SrcIp     string
	DstIp     string
	DstPort   string
	ProtoType uint32
	Action    string
	Priority  uint32
}

func (r *Rule) String() string {
	return fmt.Sprintf("(%s, %s, %s, %d, %s, %d)", r.SrcIp, r.DstIp, r.DstPort, r.ProtoType, r.Action, r.Priority)
}

type RuleList []*Rule

func (list RuleList) Len() int {
	return len(list)
}

func (list RuleList) Less(i, j int) bool {
	if list[i].Priority < list[j].Priority {
		return true
	} else {
		return false
	}
}

func (list RuleList) Swap(i, j int) {
	var temp *Rule = list[i]
	list[i] = list[j]
	list[j] = temp
}

func main() {
	rule1 := &Rule{"192.168.1.3", "", "234", 1, "ACCEPT", 30}
	rule2 := &Rule{"192.168.1.5", "", "1234", 2, "ACCEPT", 10}
	rule3 := &Rule{"0.0.0.0", "", "1000", 1, "ACCEPT", 50}
	rule4 := &Rule{"", "", "1024", 3, "ACCEPT", 30}

	ruleList := RuleList([]*Rule{rule1, rule2, rule3, rule4})
	fmt.Println(ruleList)
	sort.Sort(ruleList)
	fmt.Println(ruleList)
}
