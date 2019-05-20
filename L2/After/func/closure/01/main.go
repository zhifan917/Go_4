package main

import "fmt"

// 定义一个函数a，返回值是一个函数
func a(i string) func() {
	return func() { //该匿名函数中本身是没有定义变量i的，而是引用了它所造的环境函数a中的变量i。这个变量i就和匿名函数打包为一个整体。在闭包的有效区间内一直是生效的。
		fmt.Printf("%s is MVP", i)
	}
}

func main() {
	r := a("harden")
	r() //相当于执行了a函数内部的匿名函数
}
