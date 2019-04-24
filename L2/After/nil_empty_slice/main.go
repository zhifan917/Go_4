package main

import "fmt"

func main() {
	//空切片
	b := []int{}
	fmt.Printf("b为:%v\n", b)
	fmt.Printf("b的内存地址：%p\n", b)

	c := []string{}
	fmt.Printf("c为:%v\n", c)
	fmt.Printf("c的内存地址：%p\n", c)

	//nil切片
	var d []int
	fmt.Printf("d为:%v\n", d)
	fmt.Printf("d的内存地址：%p\n", d)

}
