package main

import (
	"fmt"
	"math"
)

/**
 *Created By miss
 *
 *At 2018/7/31
 */
type Absr interface {
	Abs() float64
} //接口定义

type Ve struct {
	x, y float64
}

func (v Ve) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)

} //实现接口方法即实现了接口
func main() {
	var a Absr
	v := Ve{3, 4}
	a = v
	fmt.Println(a.Abs())
}
