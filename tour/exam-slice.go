package main

import (
	"fmt"
	"math"
	"tour"
)

/**
 *Created By miss
 *
 *At 2018/7/30
 */

func Pic(dx, dy int) [][]uint8 {
	var res = make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		res[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			res[i][j] = uint8(float64(i) * math.Log(float64(j)))
		}
	}
	return res
}

func main() {
	r := Pic(5, 5)
	for i := range r {
		fmt.Println(r[i])
	}

}
