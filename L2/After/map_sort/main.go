package main

import (
	"fmt"
)

func main() {
	var a = make(map[string]int, 10)
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		a[key] = i
	}
	for key, value := range a {
		fmt.Printf("key:%s = %d\n", key, value)
	}
}
