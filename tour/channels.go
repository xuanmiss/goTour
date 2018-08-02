package main

import "fmt"

/**
 *Created By miss
 *
 *At 2018/7/31
 */
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum

}

func main() {
	s := []int{1, 4, 5, 7, 3, 1, 7, 3}
	c := make(chan int, 2)
	go sum(s[:len(s)/3], c)
	go sum(s[len(s)/3:len(s)/3*2], c)
	go sum(s[len(s)/3*2:], c)
	x, y, z := <-c, <-c, <-c
	fmt.Println(x, y, z, x+y+z)
}
