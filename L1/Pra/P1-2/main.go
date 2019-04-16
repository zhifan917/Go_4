package main

import (
	"fmt"
)

//练习题：将字符串"hello"逆序打印出来。实现结果"olleh"

func main() {
	s1 := "hello"
	byteArray := []byte(s1) // [h e l l o]
	length := len(byteArray)
	for i := 0; i < length/2; i++ {
		byteArray[i], byteArray[length-1-i] = byteArray[length-1-i], byteArray[i]
	}
	fmt.Println(string(byteArray))

}
