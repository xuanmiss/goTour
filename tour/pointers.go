package main

import "fmt"

/**
 *Created By miss
 *
 *At 2018/7/27
 */

func main() {
	i, j := 100, 81
	p, q := &i, &j
	fmt.Println(*p, *q)
	*p += 1
	fmt.Println(i)
	*q -= 1
	fmt.Println(*p)

}
