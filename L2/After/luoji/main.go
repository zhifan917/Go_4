package main

import "fmt"

func main() {
	//逻辑运算符
	a := true
	b := false
	//两个条件都成立才为真
	fmt.Println(a && b)
	// 两个条件有一个成立就为真   |
	fmt.Println(a || b)
	// 原来为真取非就为假,原来为假取非后就为真
	fmt.Println(!a)
	fmt.Println(!b)
}
