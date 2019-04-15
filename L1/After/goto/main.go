package main

import "fmt"

func breakDemo1() {
LICHAN:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break LICHAN
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}

func continueDemo() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}

func main() {
	// 	for i := 0; i < 10; i++ {
	// 		for j := 0; j < 10; j++ {
	// 			if j == 2 {
	// 				//设置退出标签
	// 				goto breaktag
	// 			}
	// 			fmt.Printf("%v-----%v\n", i, j)
	// 		}
	// 	}
	// 	//定义一个标签
	// breaktag:
	// 	fmt.Println("结束for循环")
	breakDemo1()
	// continueDemo()
}