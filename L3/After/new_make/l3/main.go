package main

import (
	"fmt"
)

type User struct {
	Name string
	age  int
}

func test1() {
	var p = new(int)
	*p = 1000
	fmt.Printf("p:%v address:%v\n", *p, p)

	var pUser = new(User)
	(*pUser).age = 100
	pUser.Name = "user01" //正常来说规范写应该是(*pUser).Name，但是go语言针对结构体这里做了优化（切片、map不可以，依然需要规范写），可以简化写。

	fmt.Printf("user:%v\n", *pUser)
}

func test2() {
	var p = new([]int)   //new为切片分配内存
	*p = make([]int, 10) //切片需要make为其初始化才能使用

	(*p)[0] = 100
	(*p)[2] = 100

	fmt.Printf("p:%#v\n", *p)

	var p1 = new(map[string]int)   //new为map分配内存
	*p1 = make(map[string]int, 10) //map需要make为其初始化才可以使用
	(*p1)["key"] = 100
	(*p1)["key2"] = 200

	fmt.Printf("p:%#v\n", *p1)
}

func main() {
	test1()
	test2()
}
