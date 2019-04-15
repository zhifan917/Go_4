package main

import "fmt"

// var a int
// var b bool
// var c int8

// var (
// 	m int
// // 	n string
// // )
// var name = "zhifan"
// func main() {
// 	var name string = "zhifan"
// 	age := 18
// 	_age
// }

func main() {
	s := "hello中国"
	for i := 0; i < len(s); i++ {
		// fmt.Println(s[i])
		fmt.Printf("%c\n", s[i])
	}
}
