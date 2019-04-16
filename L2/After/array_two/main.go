package main

import "fmt"

func main() {
	// var a = [3][2]int{
	// 	{1, 2},
	// 	{2, 5},
	// }
	b := [2][2]string{
		{"harden", "MVP"},
	}
	// fmt.Printf("a二维数组值为:%v\n", a)
	// fmt.Printf("b二维数组值为:%v\n", b)
	for k1, v1 := range b {
		// fmt.Println(v1)
		for k2, v2 := range v1 {
			fmt.Printf("b[%d][%d]=%s\n", k1, k2, v2)
		}
	}
}
