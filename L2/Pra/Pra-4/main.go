package main

import (
	"fmt"
	"sort"
)

// 题目：请使用内置的`sort`包对数组`var a = [...]int{3, 7, 8, 9, 1}`进行排序。

func main() {
	var a = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a[:])
	fmt.Printf("数组a排序后：%v", a)
}
