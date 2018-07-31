package main

import "fmt"

/**
 *Created By miss
 *
 *At 2018/7/27
 */

func main() {

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) //defer 函数将在外层函数返回之后返回，但是defer 函数会立即求值，多个defer会被压入栈
	}
	fmt.Println("done")
}
