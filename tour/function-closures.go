package main

import "fmt"

/**
 *Created By miss
 *
 *At 2018/7/31
 */

func addr() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}

}

func main() {
	pos, neg := addr(), addr()

	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
