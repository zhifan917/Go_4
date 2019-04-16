package main

import "fmt"

func modify(b [2]int) {
	b[1] = 1
	fmt.Printf("modify b的值为：%d\n", b)
	return
}

func main() {
	a := [2]int{1, 3}
	fmt.Printf("a的值为：%d\n", a)
	b := a
	b[0] = 5
	modify(a)
	fmt.Printf("b的值为：%d\n", b)
}
