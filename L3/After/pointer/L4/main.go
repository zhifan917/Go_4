package main

import "fmt"

// 错误空地址示例，会造成panic
//func main() {
//	var a *int
//	*a = 100
//	fmt.Printf("%d\n", *a)
//}

func main() {
	a := 25
	var b *int

	if b == nil {
		fmt.Printf("b is %v\n", b)
		b = &a
		fmt.Printf("b after initialization is %v\n", b)
	}
}
