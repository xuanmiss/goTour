package main

import (
	"fmt"
)

/**
 *Created By miss
 *
 *At 2018/7/31
 */

func fibonacci() func() int {
	a, b := 0, 1

	return func() int {
		tem := a
		a, b = b, a+b
		return tem
	}
}

func main() {
	f := fibonacci()
	fmt.Println(f())
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	fmt.Println(f(), "闭包环境没有改变")
	f = fibonacci()
	fmt.Println(f(), "新的闭包生成")

}
