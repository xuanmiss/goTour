package main

import (
	"fmt"
	"strconv"
)

/**
 *Created By miss
 *
 *At 2018/7/31
 */
type IPAddr [4]byte

func (ip IPAddr) String() string {
	var res string
	lastIndex := len(ip) - 1
	for i, v := range ip {
		res += strconv.Itoa(int(v))
		if i != lastIndex {
			res += "."
		}
	}
	return res
}

func main() {
	a := IPAddr{1, 2, 3, 4}
	b := IPAddr{127, 0, 0, 1}
	fmt.Println(a)
	fmt.Println(b)
}
