package main

//编写一个函数接受一个字符串参数，输出该字符串是否是回文（从前往后和从后往前都是一样的字符串成为回文）

import "fmt"

func main() {
	funcation()
}

func funcation() {
	var a string
	fmt.Print("请输入字符串参数a:")
	fmt.Scanln(&a)
	// fmt.Println(a)
	//打印a的回文
	// b := ([]rune)(a)
	b := []rune(a)
	for i := 0; i < len(b)/2; i++ {
		c := b[i]
		b[i] = b[len(b)-i-1]
		b[len(b)-i-1] = c
		// fmt.Println(c, b[i], b[len(b)-i-1])
	}
	d := string(b)
	fmt.Println(string(d))

}
