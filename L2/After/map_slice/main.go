package main

import "fmt"

func main() {
	//定义元素为map类型的切片
	mapSlice := make([]map[string]string, 3)

	for k, v := range mapSlice {
		fmt.Printf("index:%d value：%v\n", k, v)
	}
	fmt.Println("after init!!!")

	//对切片中map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "哈登"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "休斯顿"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}
