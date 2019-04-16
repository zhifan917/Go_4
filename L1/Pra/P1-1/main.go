package main

import "fmt"

//将字符串"hello"逆序打印出来。实现结果"olleh"

func main() {
	s1 := "hello"
	byteArray := []byte(s1) //转化为字节数组进行操作 [h e l l o]
	s2 := ""                //定义一个空字符串进行接收
	for i := len(byteArray) - 1; i >= 0; i-- {
		// i 是 4 3 2 1 0
		// byteArray[i]  o l l e h （字符）
		fmt.Println(i)
		s2 = s2 + string(byteArray[i])
	}
	fmt.Println(s2)
}
