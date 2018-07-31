package main

import (
	"fmt"
	"math"
)

/**
 *Created By miss
 *<p> go for and func exam
 *At 2018/7/27
 */

func mySqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(i, z)
	}
	return z
}

func main() {
	fmt.Println("my fun", mySqrt(2))
	fmt.Println("math's", math.Sqrt(2))

}
