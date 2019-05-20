package main

import "fmt"

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}
func main() {
	var f = adder() //变量f就是一个闭包。此时f就相当于是一个func(int)int 类型的函数
	fmt.Println(f(10)) //10  x=0 y=10 
	fmt.Println(f(20)) //30  x=10 y=20
	fmt.Println(f(30)) //60  x=30 y=30

	f1 := adder()  //变量f1就是一个闭包
	fmt.Println(f1(40)) //40 x=0 y=40
	fmt.Println(f1(50)) //90 x=40 y=50
}
