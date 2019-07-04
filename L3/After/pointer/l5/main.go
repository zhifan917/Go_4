package main

import "fmt"

// 操作指针变量指向的地址里面的值
func pointer1() {
	b := 255
	a := &b
	fmt.Println("address of b is", &b)
	fmt.Println("address of a is", a)
	fmt.Println("value of b is", *a) //获取指针变量中存的内存地址对应的值

	*a = 90 //修改指针变量的存的内存地址所对应的值
	fmt.Println("address of b is", b)
}

// 通过指针修改变量的值
func pointer2() {
	b := 255
	a := &b
	fmt.Println("address of b is", &b)
	fmt.Println("address of a is", a)
	fmt.Println("value of b is", *a)
	*a++ //修改指针变量的存的内存地址所对应的值
	fmt.Println("new value of b is", b)
}
func main() {
	//pointer1()
	pointer2()
}
