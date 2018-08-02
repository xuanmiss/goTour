package main

import (
	"fmt"
	"time"
)

/**
 *Created By miss
 *
 *At 2018/7/31
 */

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(2000 * time.Millisecond)
		fmt.Println(s)
	}

}

func main() {
	go say("hello")
	say("world")

}
