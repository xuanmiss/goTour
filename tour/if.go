package main

import (
	"fmt"
	"math"
	"strings"
)

/**
 *Created By miss
 *
 *At 2018/7/27
 */
func ifgo(x string) string {
	var res string
	if strings.Count(x, "")-1 < 10 {
		res = "这是一个短字符串"
	} else {
		res = "这是一个长字符串"
	}
	fmt.Println(strings.Count(x, "") - 1)
	return res
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim

}

func main() {

	fmt.Println(ifgo("这里一共有10个字符"))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
