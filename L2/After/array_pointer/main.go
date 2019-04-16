package main

import "fmt"

func main() {
	x, y := 1, 2
	var arr = [...]int{5: 2}
	//数组指针
	var pf *[6]int = &arr

	//指针数组
	pfArr := [...]*int{&x, &y}
	fmt.Println(pf)        //此处pf就表示*pf 这是go编译器会推断pf的类型，发现pf是个指针类型，在调用时会变成*pf
	fmt.Printf("%T\n", pf) //打印pf的变量类型
	fmt.Println(pfArr)
}
