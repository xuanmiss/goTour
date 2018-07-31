package main

import (
	"fmt"
	"math"
)

func add(x int, y int) (int, int) {

	return x + y, x
}

/*
 *func 可以接受任意数个参数，返回任意数个返回值
 *func 参数类型一致可统一声明类型
 */
func swap(x, y string) (string, string) {

	return y, x
}

func main() {

	sum, a := add(3, 6)
	fmt.Println("sum is ", sum, "\nx is ", a) //Println()用,来连接字符串或数据

	fmt.Println(swap("world", "hello"))
	//fmt.Printf("Now you have %g problems.",math.Sqrt(7))
	fmt.Println(math.Pi)
}
