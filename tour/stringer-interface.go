package main

import (
	"fmt"
)

/**
 *Created By miss
 *
 *At 2018/7/31
 */

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}
func main() {
	a := Person{"zhangsan", 18}
	b := Person{"lisi", 15}
	fmt.Println(a, b)
}
