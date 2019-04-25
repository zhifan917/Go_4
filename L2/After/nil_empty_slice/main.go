package main

import "fmt"

func main() {
	//空切片
	b := []int{}
	fmt.Printf("b为:%#v\n", b) //%#v表示以go语言语法形式显示值
	fmt.Printf("b的内存地址：%p\n", b)
	b = append(b, 2)
	fmt.Printf("After:b为:%#v\n", b)
	fmt.Printf("After:b的内存地址：%p\n", b)

	c := []string{}
	fmt.Printf("c为:%#v\n", c)
	fmt.Printf("c的内存地址：%p\n", c)
	c = append(c, "MVP")
	fmt.Printf("After:c为:%#v\n", c)
	fmt.Printf("After:c的内存地址：%p\n", c)

	//nil切片
	var d []int
	fmt.Printf("d为:%#v\n", d)
	fmt.Printf("d的内存地址：%p\n", d)
	d = append(d, 1)
	fmt.Printf("After:d为:%#v\n", d)
	fmt.Printf("After:d的内存地址：%p\n", d)

}
