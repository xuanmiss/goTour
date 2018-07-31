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

type Vertex struct {
	x, y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
