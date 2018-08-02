package main

import (
	"fmt"
	"time"
)

/**
 *Created By miss
 *error判断函数执行结果
 *At 2018/7/31
 */

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}
func run(i int) error {
	var e error
	if i > 10 {
		e = &MyError{
			time.Now(),
			"it didn't work",
		}
	} else {
		e = nil
	}
	return e

}

func main() {
	if err := run(11); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
}
