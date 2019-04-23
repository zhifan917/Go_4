package main

import "fmt"

func main() {
	a := [3]string{"a", "b", "c"}
	b := a[1:]
	fmt.Println(b)
}
