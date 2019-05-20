package main

import (
	"fmt"
	"strings"
)

func makeSuffixFunc(suffix string) func(string) string { //suffix是全局变量
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) { //string.HasSuffix(name,suffix) 表达的是：name以suffix为结尾返回true
			return name + suffix
		}
		return name
	}
}
func main() {
	func1 := makeSuffixFunc(".bmp") //suffix=.bmp
	func2 := makeSuffixFunc(".jpg") //suffx=.jpg
	fmt.Println(func1("test"))      //name=test
	fmt.Println(func2("test"))      //name=test
}
