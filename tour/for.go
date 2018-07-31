package main

import "fmt"

/**
 *Created By miss
 *
 *At 2018/7/27
 */

func main() {

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sm := 1
	for sm < 100 {
		fmt.Println("hello")
	} //相当于while
	/*	for {
			fmt.Println()
		}无限循环 = while true
	*/
	//fmt.Println(sm)
}
