package main

import (
	"fmt"
	"strings"
)

/*
题目：
写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
*/
func main() {
	//0.定义一个map[string]int
	var s = "how do you do"
	var wordCount = make(map[string]int, 10)
	//1.字符中都有哪些单词
	words := strings.Split(s, " ")
	//2.遍历单词做统计
	for _, word := range words {
		v, ok := wordCount[word]
		if ok {
			//map中有这个单词记录
			wordCount[word] = v + 1
		} else {
			wordCount[word] = 1
		}
	}
	for k, v := range wordCount {
		fmt.Printf("%s=%d\n", k, v)
	}
}
