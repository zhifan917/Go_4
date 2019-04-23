package main

import "fmt"

func copySlice() {
	a := []int{1, 2, 3}
	c := make([]int, 3, 4)
	copy(c, a)
	fmt.Printf("c:%d\n", c)
	c[0] = 100
	fmt.Printf("c:%d\n", c)
	fmt.Printf("a:%d\n", a)
}

func main() {
	copySlice()
}
