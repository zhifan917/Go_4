package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{
		"harden": 100,
		"kobe":   80,
		"james":  70,
		"wade":   200,
		"cury":   180,
	}
	// 题目： 将m按照key的ASCII顺序打印键值对
	// 1. 先把key取出来放到切片里面
	var keys = make([]string, 0, 10)
	for key := range m {
		keys = append(keys, key)
	}
	fmt.Println(keys)
	// 2. 对key的切片做排序
	sort.Strings(keys)
	fmt.Println(keys)
	// 3. 按照排序后的key依次去map取值
	for _, key := range keys {
		fmt.Println(key, m[key])
	}
}
