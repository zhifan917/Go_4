package main

import (
	"fmt"
)

func main() {
	s1 := "hello"
	byteArray := []byte(s1) // [h e l l o]
	length := len(byteArray)
	for i := 0; i < length/2; i++ {
		byteArray[i], byteArray[length-1-i] = byteArray[length-1-i], byteArray[i]
	}
	fmt.Println(string(byteArray))

}
