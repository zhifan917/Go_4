package main

import (
	"fmt"
	"sort"
)

func main() {
	key := []int{1, 5, 2, 5, 2}
	sort.Ints(key)
	fmt.Println(key)
}
