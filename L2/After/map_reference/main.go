package main

import "fmt"

func main() {
	a := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}

	a["mike"] = 9000
	fmt.Println("origin map", a)
	b := a
	b["mike"] = 18000
	fmt.Println("a map changed", a)
}
