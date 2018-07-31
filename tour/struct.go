package main

import (
	"fmt"
)

/**
 *Created By miss
 *
 *At 2018/7/27
 */

//type man struct{
//	name string
//	email string
//}

var (
	wangwu = man{email: "wangwu@gmail.com"}
)

func main() {
	fmt.Println(man{"miss", "miss@gmail.com"})
	lisi := man{"lisi", "lisi@gmail.com"}

	fmt.Println(lisi.email)
	p := &lisi
	p.email = "xiugai.gmail.com"
	fmt.Println(lisi.email)
	fmt.Println(*p, p)

	fmt.Println(wangwu)

}
