package main

import "fmt"

func calc(a, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}
func main() {
	sum, sub := calc(2, 3)
	fmt.Println(sum, sub)
}
