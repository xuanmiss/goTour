package main

import (
	"fmt"
	"strings"
)

/**
 *Created By miss
 *
 *At 2018/7/30
 */

func WordCount(ss string) map[string]int {
	res := make(map[string]int)
	s := strings.Fields(ss)
	for i := 0; i < len(s); i++ {
		_, ok := res[s[i]]
		if ok == true {
			res[s[i]] += 1 //存在+1
		} else {
			res[s[i]] = 1 //不存在 赋值
		}

	}
	return res
}

func main() {
	fmt.Println(WordCount("hello , i am a hello a"))
}
