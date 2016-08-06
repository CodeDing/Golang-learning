package main

import (
	"fmt"
	"sort"
)

//对任意对象进行排序
type Person struct {
	name    string
	company string
	age     uint32
}

//为*Person添加String()方法，便于输出
func (p *Person) String() string {
	return fmt.Sprintf("( %s, %s, %d )", p.name, p.company, p.age)
}

type PersonList []*Person

//排序规则：首先按年龄排序（由小到大），年龄相同时按姓名进行排序（按字符串的自然顺序）

func (list PersonList) Len() int {
	return len(list)
}

func (list PersonList) Less(i, j int) bool {
	if list[i].age < list[j].age {
		return true
	} else {
		return false
	}
}

func (list PersonList) Swap(i, j int) {
	var temp *Person = list[i]
	list[i] = list[j]
	list[j] = temp
}

func main() {
	fmt.Println("------")
	p1 := &Person{"Tom", "intel", 19}
	p2 := &Person{"Hanks", "ibm", 20}
	p3 := &Person{"Amy", "alibaba", 19}
	p4 := &Person{"Tom", "ge", 20}
	p5 := &Person{"Jogn", "quatum", 21}
	p6 := &Person{"Mike", "ms", 23}

	pList := PersonList([]*Person{p2, p1, p4, p5, p3, p6})
	fmt.Println(pList)
	sort.Sort(pList)
	fmt.Println(pList)

}
