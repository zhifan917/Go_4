package main

import "fmt"

func main() {
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}
