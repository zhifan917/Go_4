package main

import "fmt"

func main() {
	// //空切片
	// b := []int{}
	// b[0] = 1
	// fmt.Printf("b为:%v\n", b)
	// fmt.Printf("b的内存地址：%p\n", b)

	//nil切片
	var d []int
	d[0] = 1
	fmt.Printf("d为:%v\n", d)
	fmt.Printf("d的内存地址：%p\n", d)

}
