package main

import (
	"fmt"
)

func adder() func(int) int {
	var x int
	return func(d int) int {
		x = x + d
		return x
	}
}

func main() {
	f := adder()
	fmt.Println(f(100))
	fmt.Println(f(200))

	f1 := adder() //Adder函数重新做了实例化
	fmt.Println(f1(1))
	fmt.Println(f1(2))
}
