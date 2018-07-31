package main

import (
	"fmt"
)

/**
 *Created By miss
 *
 *At 2018/7/30
 */

func arr() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a)
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	fmt.Println("len a is ", len(a), "cap a is ", cap(a))
}

func printcap(a []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)
}

func main() {
	arr()

	type man struct {
		name  string
		email string
	}

	primes := [6]int{2, 3, 5, 7, 11, 13} //声明数组primes
	var s []int = primes[0:6]            //创建数组primes的切片s
	var s1 []int = primes[:2]
	fmt.Println(s, s1)
	s[0] = 100 //切片操作会改变数组的值和其他引用切片
	fmt.Println("hello", "world")
	fmt.Println(s, s1, primes)

	m := []man{
		{"lisi", "lisi@gmail.com"},
		{"zhangsan", "zhangsan@gamil.com"},
	}
	fmt.Println(m)
	s = s[:0]
	printcap(s)

	s = s[:4]
	printcap(s)

	s = s[2:]
	printcap(s) //cap是从切片的起始元素开始到底层数组最后一元素的长度

}
