package main

import "fmt"

func main() {
	// //方法1：
	// var a [3]int
	// a[0] = 1
	// a[1] = 2
	// a[2] = 3
	// fmt.Printf("数组的值为%v\n", a)

	// //方法2：
	// var a [2]int = [2]int{1, 5}
	// fmt.Printf("数组的值为%v\n", a)

	//方法3：
	// a := [2]int{1, 7}
	// //方法4：
	// b := [...]int{3, 9, 10}
	// //方法5:
	// c := [3]int{8}
	// //方法6：
	// d := [...]int{1: 4, 6: 9}
	// fmt.Printf("a数组的值为%v\n", a)
	// fmt.Printf("b数组的值为%v\n", b)
	// fmt.Printf("c数组的值为%v\n", c)
	// fmt.Printf("d数组的值为%v\n", d)

	//练习
	var city = [5]string{"北京", "上海"}
	star := [...]string{"harden", 4: "james"}
	fmt.Printf("城市有%s\n", city)
	fmt.Printf("球星有%s\n", star)
}
