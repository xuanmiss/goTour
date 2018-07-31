package main

import "fmt"

/**
 *Created By miss
 *
 *At 2018/7/30
 */

type user struct {
	name string
	age  int
}

var m map[string]user

func main() {
	m = make(map[string]user)
	m["zhangsan"] = user{"zhangsan", 19}
	fmt.Println(m["zhangsan"])

	var mm = map[string]user{
		"lisi":   {"lisi", 17},
		"wangwu": {"wangwu", 18},
	} //map 一组映射关系
	for i, v := range mm {
		fmt.Println(i, v) //可用range来循环
	}

	mmm := make(map[string]int)
	mmm["one"] = 1
	fmt.Println(mmm["one"])
	mmm["one"] = 3
	fmt.Println(mmm["one"])
	delete(mmm, "one")
	elm, ok := mmm["one"]
	if ok == true {
		fmt.Println(elm, ok)
	} else {
		fmt.Println("不存在")
	}

}
