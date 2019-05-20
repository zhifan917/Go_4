package main

import (
	"fmt"
)

// Add 可变参数函数
func Add(a ...int) int { //边长参数a，需要传入0个或多个参数
	fmt.Printf("func args count:%d\n", len(a)) //变长参数在底层存储就是一个切片，计算切片的长度，其实就是计算传入参数的个数
	var sum int
	for index, args := range a { //遍历切片，打印出传入参数值
		fmt.Printf("args[%d]=%d\n", index, args)
		sum = sum + args //求和
	}
	return sum
}

func testAdd() {
	sum := Add() //传入0个参数
	fmt.Printf("sum=%d\n", sum)

	sum = Add(2) //传入1个参数
	fmt.Printf("sum=%d\n", sum)

	sum = Add(5, 98) //传入2个参数
	fmt.Printf("sum=%d\n", sum)
}
func main() {
	testAdd()
}
