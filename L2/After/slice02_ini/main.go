package main

import (
	"fmt"
)

func empty() {
	var a []string
	if a == nil {
		fmt.Println("a是空切片")
	}
	a[0] = "paul" //不初始化直接赋值操作
	fmt.Println(a)
}

func emptyAppend() {
	var a []string
	if a == nil {
		fmt.Println("a是空切片")
	}
	a = append(a, "harden") //不初始化但是可以进行append操作，append会自动对这个空切片做自动扩容。
	fmt.Println(a)
}

func sliceArray() {
	//基于数组定义切片
	a := [5]int{1, 2, 3, 4, 5}
	b := a[1:4] //基于数组a创建切片，包括a[1] a[2] a[3]
	fmt.Println(b)
	fmt.Printf("b的类型为：%T\n", b)
}

func sliceInSlice() {
	//切片在切片
	a := [...]string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
	b := a[1:3]
	fmt.Printf("b:%v,type:%T,len:%d,cap:%d\n", b, b, len(b), cap(b))
	c := b[1:]
	fmt.Printf("c:%v,type:%T,len:%d,cap:%d\n", c, c, len(c), cap(c))
}

func sliceMake() {
	a := make([]bool, 2, 5)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
func main() {
	sliceMake()
}
