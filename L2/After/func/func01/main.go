package main

import "fmt"

func intSum(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

func intSum2(x int, y ...int) int {
	fmt.Println(x, y) //是一个切片
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}
func main() {
	// ret1 := intSum()
	// ret2 := intSum(10)
	// ret3 := intSum(10, 20)
	// ret4 := intSum(10, 20, 30)
	// fmt.Println(ret1, ret2, ret3, ret4)
	// ret1 := intSum2()
	// ret2 := intSum2(10)
	// ret3 := intSum2(10, 20)
	// ret4 := intSum2(10, 20, 30)
	// fmt.Println(ret1, ret2, ret3, ret4)
}
