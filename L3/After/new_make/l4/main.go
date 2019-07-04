package main

import "fmt"

func value() {
	var a int = 100
	fmt.Printf("a addr is %p\n", &a)
	b := a
	fmt.Printf("b addr is %p\n", &b)
	a = 50
	fmt.Printf("a addr is %p\n", &a)
	fmt.Printf("a=%d,b=%d", a, b)
}

func pointer() {
	var a = 100
	var b = &a
	var c = b
	*c = 200
	fmt.Printf("a=%v b=%v c=%v", a, *b, *c)
}

// test
func main() {
	value()
	pointer()
}
