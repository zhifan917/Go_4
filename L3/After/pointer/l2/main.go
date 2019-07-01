package main

import "fmt"

func main() {
	var b int32
	b = 156
	var a *int32
	fmt.Printf("addr of a:%v\n", a)

	a = &b
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", *a)
}
