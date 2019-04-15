package main

import (
	"fmt"
	
)

func main() {
	var a int = 10
	fmt.Printf("%d \n", a)
	fmt.Printf("%b \n", a)

	var b int = 077
	fmt.Printf("%o \n", b) // 77


	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF

	// 变量的内存地址
	fmt.Printf("%p \n", &a) // 0xc00004c080  占位符%p表示十六进制的内存地址
}
