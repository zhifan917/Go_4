package main

import "fmt"

func main() {
	//练习题：求数组`[1, 3, 5, 7, 8]`所有元素的和
	a := [...]int{1, 3, 4, 7, 8}
	sum := 0
	for _, v := range a {
		sum = sum + v
	}
	fmt.Printf("数组的和为%d\n", sum)
}
