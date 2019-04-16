package main

import "fmt"

func main() {
	a := [5]int{1, 9, 17}
	// //方法1：
	// for i := 0; i < len(a); i++ {
	// 	fmt.Printf("a数组的值为%v\n", a[i])
	// }

	//方法2： for range
	for k, v := range a {
		fmt.Printf("index:%d, value:%d\n", k, v)
	}
}
