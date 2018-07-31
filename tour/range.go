package main

import (
	"fmt"
)

/**
 *Created By miss
 *
 *At 2018/7/30
 */

func main() {
	var pow = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range pow {
		fmt.Println(i, "*", v, "=", i*v)
	}

	var p = make([]int, 10)
	for i := range p {
		p[i] = i << uint(i)
	}

	for _, v := range p {
		fmt.Println(v)
	}

}
