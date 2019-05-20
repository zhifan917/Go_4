package main

import (
	"fmt"
)

func calc(op func(args ...int) int, opArgs ...int) int { //第一个参数op是一个函数类型的参数，函数是一个接收可变类型参数，有一个返回值；第二个参数直接是一个可变参数
	result := op(opArgs...) //将第2个参数opArgs传给op函数，并且是切片，要传递需要将切片展开，也就是切片名...
	fmt.Printf("result = %d\n", result)
	return result
}

func add(args ...int) int {
	var sum int
	for i := 0; i < len(args); i++ {
		sum = sum - args[i]
		fmt.Println(i)
	}
	return sum
}

func main() {
	calc(func(args ...int) int { //因为参数类型是匿名函数，所以我们此处调用传参，也需要传入匿名函数
		//return 0 //1就是op的返回值
		var sum int
		for i := 0; i < len(args); i++ { //这里args的参数其实就是calc函数中第二个可变参数的值（因为上面op函数已经将可变（切片）参数值展开了）
			fmt.Printf("sum=%d\n", sum)
			sum = sum + args[i]
			fmt.Println(i)
		}
		return sum
	}, 1, 2, 3, 4, 5)

	calc(func(args ...int) int { //因为参数类型是匿名函数，所以我们此处调用传参，也需要传入匿名函数
		//return 0 //1就是op的返回值
		var sum int
		for i := 0; i < len(args); i++ {
			sum = sum - args[i]
		}
		return sum
	}, 1, 2, 3, 4, 5)

	calc(add, 1, 3) //不单单是匿名函数可以传参，有名函数传的值、类型和给的返回值是同一类型，其就是同一类型函数
}
