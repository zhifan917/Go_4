package main

import "fmt"

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

func main() {
	//// 指针取值
	//a := 10
	//b := &a
	//fmt.Printf("type of b:%T\n", b)
	//
	//c := *b
	//fmt.Printf("type of c:%T\n", c)
	//fmt.Printf("value of c:%v\n", c)
	a := 10
	modify1(a)
	fmt.Println(a) // 10
	modify2(&a)
	fmt.Println(a) // 100
}
