package main

import "fmt"

// 加密
func main() {
	var a *int
	var b = 200

	a = &b
	fmt.Printf("value of a %v\n", a)
	fmt.Printf("address of b %v\n", &b)
	fmt.Printf("address of a %v\n", &a)

	*a = 300
	fmt.Printf("value of b %v\n", b)
	fmt.Printf("type of a %T\n", a)

}
