package main

import "fmt"

func main() {
	scoreMap := map[string]int{
		"张三": 90,
		"小明": 100,
		"张飞": 80,
	}
	for k := range scoreMap {
		fmt.Println(k)
	}
}
