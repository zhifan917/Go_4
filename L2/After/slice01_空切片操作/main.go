package main

import "fmt"

func main() {
	// // 声明切片类型
	var a []string //声明一个字符串切片
	// fmt.Println(len(a))
	// fmt.Println(cap(a))

	if a == nil {
		fmt.Println("a为空切片")
	}
	fmt.Println(a)
	a = append(a, "harden")
	fmt.Println(a)
	// //操作空切片
	// a[0] = "harden"
	// fmt.Println(a)

	// var b = []int{}
	// // b[0] = 2
	// fmt.Println(b)
	// a := make([]int, 2, 10)
	// a[0] = 1
	// fmt.Println(a)
	b := []int{}

	if b == nil {
		fmt.Println("b为空切片")
	}
	fmt.Println(b)
	b = append(b, 1)
	fmt.Println(b)

}
