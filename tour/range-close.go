package main

import "fmt"

/**
 *Created By miss
 *
 *At 2018/7/31
 */
func fibon(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func num(c chan int) {

	for i := 0; i < 10; i++ {
		if i > 5 {
			c <- i
		} else {
			c <- 10 + i
		}
	}
	close(c)
}

func main() {
	c := make(chan int, 10) //缓冲大小
	go fibon(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
	cc := make(chan int, 10)
	go num(cc)
	for i := range cc {
		fmt.Println(i)
	}
}
