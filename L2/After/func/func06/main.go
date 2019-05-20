package main

import "fmt"

type functionType func(int) bool // 声明了一个函数类型

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}
func filter(slice []int, f functionType) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}
func main() {
	f1 := filter                      //将函数add赋值给变量f1
	fmt.Printf("type of f1:%T\n", f1) //type of f1:func(int, int) int
	slice := []int{1, 2, 3, 4, 5, 7}
	ret := f1(slice, isOdd)
	fmt.Println(ret)
}
