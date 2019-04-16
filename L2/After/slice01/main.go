package main

import "fmt"

func main() {
	// 声明切片类型
	var a []string //声明一个字符串切片
	fmt.Println(len(a))
	fmt.Println(cap(a))
	if a == nil {
		fmt.Println("a为空切片")
	}

}
