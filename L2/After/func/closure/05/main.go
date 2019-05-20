package main

import "fmt"

func add(base int) func(int) int { //base在闭包中是全局变量
	return func(i int) int {
		base += i
		return base
	}
}
func main() {
	tmp1 := add(10)               //base=10
	fmt.Println(tmp1(1), tmp1(2)) //i=1 base=10+1=11 i=2 base=11+2=13
	tmp2 := add(100)              //base=100
	fmt.Println(tmp2(1), tmp2(2)) //i=1 base=100+1=101 i=2 base=101+2=103
}
