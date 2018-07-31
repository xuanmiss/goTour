package main

import "fmt"

/**
 *Created By miss
 *
 *At 2018/7/30
 */
func printsl(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func main() {

	var s []int
	printsl(s)

	// append works on nil slices.
	s = append(s, 0)
	printsl(s)

	// The slice grows as needed.
	s = append(s, 1)
	printsl(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4, 5, 6, 7, 8)
	printsl(s)
}
