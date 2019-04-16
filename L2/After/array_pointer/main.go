package main

import "fmt"

func main() {
	x, y := 1, 2
	var arr = [...]int{5: 2}
	//数组指针
	// var pf *[6]int = &arr
	var pf = &arr
	//等价于var pf *[6]int = &arr 小知识点：1、go语言不推荐我们在左边重复在写变量的类型，因为编辑器会根据右边的值来自动匹配确定；
	//2、长度也是数组类型的一部分，所以等号左边数组长度是多少，右边也是多少，想要数组赋值，两个数组就必须是同一个类型才行。

	//指针数组
	pfArr := [...]*int{&x, &y}
	fmt.Println(pf)        //此处pf就表示*pf 这是go编译器会推断pf的类型，发现pf是个指针类型，在调用时会变成*pf
	fmt.Printf("%T\n", pf) //打印pf的变量类型
	fmt.Println(pfArr)
}
