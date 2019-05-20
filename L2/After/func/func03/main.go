package main

import "fmt"

func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}
func main() {
	sum, sub := calc(2, 3)
	fmt.Println(sum, sub)
}
