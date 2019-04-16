package main

import "fmt"

func main() {
	type Stu struct {
		Name string
	}
	pf := &Stu{ // pf是指针
		Name: "xxx",
	}
	fmt.Println(pf.Name) // 这是go编译器会推断pf的类型，发现pf是个指针类型，在调用时会变成(*pf).Name
	fmt.Println((*pf).Name)
	fmt.Println(pf)
	ll := (&(*pf).Name)
	fmt.Println(ll)
}
