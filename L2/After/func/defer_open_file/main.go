package main

import (
	"fmt"
	"os"
)

func testDefer() {
	/*
		os也是go语言标准包（用于操作系统），打开一个文件会返回2个值，第一个是文件打开成功，会返回一个文件的对象，
		通过这个文件对象，就可以去操纵这个文件了（打开、修改等），第二个返回的是错误信息
	*/
	file, err := os.Open("C:/tmp.txt")
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}

	/*
		这里我们使用defer语句，因为defer语句是在函数退出时最后执行，所以当我们文件对象成功生成后，在这里加上该条语句，
		那么不论下面哪一条语句成功执行要退出函数时，都会在退出函数时执行该语句，这样就不需要我们去在每一条语句写file.close了。
	*/
	defer file.Close()

	var buf [4096]byte
	n, err := file.Read(buf[:]) //要传的是切片,如果是数组，值类型改不了外面真正的值，改的是副本。
	if err != nil {
		fmt.Printf("read file failed,err:%v\n", err)
		//file.Close()  //文件成功打开后需要关闭,如果没有defer的话，此处就需要写close关闭文件
		return
	}
	fmt.Printf("read %d byte succ, content:%v\n", n, string(buf[:]))
	//file.Close()    //文件成功打开后需要关闭，如果没有defer的话，此处就需要写close关闭文件
	return
}
func main() {
	testDefer()
}
