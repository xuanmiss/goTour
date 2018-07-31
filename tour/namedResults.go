package main

import "fmt"

/**
 *Created By miss
 *
 *At 2018/7/27
 */

func split(sum int) (y, x int) {

	x = sum % 10
	y = sum - x
	return
}

func main() {

	fmt.Print(split(17))

}
