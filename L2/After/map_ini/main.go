package main

import "fmt"

func mapMake() {
	// var user map[string]int = make(map[string]int,5000)
	user := make(map[string]int, 5000)
	user["abc"] = 38
	fmt.Println(user)
}

func mapNormal() {
	m := map[string]int{"张三": 21}
	m1 := map[string]int{"张三": 21, "李四": 22}
	m2 := map[string]int{
		"张三": 21,
		"李四": 22,
	}
	m3 := map[string]int{}
	fmt.Println(m)
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
}

func mapNormal2() {
	a := map[string]int{
		"harden": 100,
		"james":  90,
	}
	a["james"] = 1
	fmt.Println(a)
}
func main() {
	// mapMake()
	// mapNormal()
	mapNormal2()
}
