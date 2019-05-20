package main

import "fmt"

func calc(base int) (func(int) int, func(int) int) { //base是全局变量 ，返回值是2个匿名函数
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
func main() {
	f1, f2 := calc(10)        //f1和f2引用的是同一个base变量（base=10），相当于base变量在add和sub两个匿名函数内都是生效的，是同一个。
	fmt.Println(f1(1), f2(2)) //f1：base=10 i=1 base=10+1=11 f2:base=11 i=2 base=11-2=9
	fmt.Println(f1(3), f2(4)) //f1：base=9 i=3 base=9+3=12 f2:base=12 i=4 base=12-4=8
	fmt.Println(f1(5), f2(6))
	fmt.Println(f1(7), f2(8))
}
