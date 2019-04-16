package main

import "fmt"

func main() {
	//练习2： 找出数组中和为指定值的两个元素的下标，比如从数组`[1, 3, 5, 7, 8]`中找出和为8的两个元素的索引下标。  结果：索引分别为`(0,3)`和`(1,2)`。
	/*
			思路：遍历数组
			2.1依次取出每个元素
			2.2 计算一下 other= 8-当前值
		    2.3 在不在数组中，在的话把索引拿出来
	*/
	a1 := [...]int{1, 3, 5, 7, 8}
	for index, value := range a1 {
		other := 8 - value
		for index2 := index + 1; index2 < len(a1); index2++ { //避免索引重复
			if a1[index2] == other {
				// 另一半在数组中，把它们的索引打印出来
				fmt.Printf("和为8的索引是：(%d %d)\n", index, index2)
			}
		}
	}

}
