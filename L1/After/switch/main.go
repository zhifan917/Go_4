package main

import "fmt"

func switch1() {
	i := 5
	switch i {
	case 1:
		fmt.Println("小拇指")
	case 5:
		fmt.Println("大拇指")
	case 8:
		fmt.Println("无名指")
	default:
		fmt.Println("无效的输入")
	}
}

func switch2() {
	switch i := 5; i {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println("无效的输入")
	}
}

func switch3() {
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
}

func switch4() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	default:
		fmt.Println(".....")
	}

}
func main() {
	// switch1()
	// switch3()
	switch4()
}
