package main

import "fmt"

func main() {
	// var (
	// 	a = 100
	// 	// b = "henry"
	// )

	// // %v俗称占位符
	// fmt.Printf("a=%v\n", a)
	// // %T 打印类型
	// fmt.Printf("b的类型是%T\n", b)
	// // %%转义
	// fmt.Printf("%d%%\n", a)

	// fmt.Printf("|%d|\n", a)
	// fmt.Printf("|%8d|\n", a)
	// fmt.Printf("|%-8d|\n", a)
	// fmt.Printf("|%08d|\n", a)
	// f1 := 3.141592654
	// fmt.Printf("%.2f\n", f1) //最多2位小数表示
	// fmt.Printf("%.2g\n", f1) //最多用2位数字表示
	s1 := "这是一个字符串\""
	// fmt.Printf("s1:%s\n", s1) // %s 正常输出字符串
	// fmt.Printf("s1:%q\n", s1) // %q 字符串带双引号，字符串中的引号带转义符

	fmt.Printf("s1:%20s\n", s1) //字符串最小宽度为20
	fmt.Printf("s1:%.5s\n", s1) //字符串最大宽度为5
}
