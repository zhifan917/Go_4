package main

import "fmt"

//练习题：打印9*9乘法表
func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		//换行
		fmt.Println()
	}
}
