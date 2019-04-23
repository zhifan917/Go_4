package main

import "fmt"

func sliceDel() {
	a := []int{171, 232, 11, 44, 555, 222, 111}

	fmt.Printf("删除前：%d\n", a)
	//删除索引为3的元素
	a = append(a[:3], a[4:]...)
	fmt.Printf("删除后：%d\n", a)
}

func main() {
	sliceDel()
}
