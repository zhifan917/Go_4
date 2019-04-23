package main

import "fmt"

func sliceAppend() {
	var numslice []int
	for i := 0; i < 10; i++ {
		numslice = append(numslice, i)
		fmt.Printf("numslice:%v ,len:%d ,cap:%d ,ptr:%p\n", numslice, len(numslice), cap(numslice), numslice)
	}
}

func appendMany() {
	var cityslice []string
	cityslice = append(cityslice, "beijing") //追加1个元素
	fmt.Println(cityslice)

	cityslice = append(cityslice, "shanghai", "shenzhen", "taibei") //追加多个元素
	fmt.Println(cityslice)

	//追加切片
	a := []string{"yuncheng", "taiyuan"}
	cityslice = append(cityslice, a...)
	fmt.Println(cityslice)
}
func main() {
	appendMany()
}
