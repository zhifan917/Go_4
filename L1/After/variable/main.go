// var name string = "henry"
// var age int = 18

// var name,age = "henry"
package main

import "fmt"

func foo() (int, string) {
	return 10, "henry"
}

func main() {
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}
