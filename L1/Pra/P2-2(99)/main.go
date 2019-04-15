package main

import "fmt"

//反着打印9*9乘法表
func main() {
	for i := 1; i < 10; i++ {
		for j := i; j < 10; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		//换行
		fmt.Println()
	}
}
