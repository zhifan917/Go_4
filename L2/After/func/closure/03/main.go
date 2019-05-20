package main

import "fmt"

func adder() func(int) int {
	var x int //x未赋值，默认为0
	return func(d int) int {
		x += d //匿名函数引用外部函数的x变量，x就相当于匿名函数中的全局变量了，不在是Adder函数中的局部变量了。
		return x
	}
}

func main() {
	var f = adder()       //f就是一个闭包实例
	fmt.Print(f(1), "-")  // x=0 d=1 x=0+1=1
	fmt.Print(f(20), "-") //x=1 d=20 x=1+20=21
	fmt.Print(f(300))     //x=21 d=300 x=21+300=321
}
