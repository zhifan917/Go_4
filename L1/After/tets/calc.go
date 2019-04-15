package calc

import (
	"fmt"
)

// Add 用于计算包
func Add(a int, b int) int {
	var c = a + b
	fmt.Printf("add result:%d\n", c)
	return c
}
