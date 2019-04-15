package main

import "fmt"

func fordemo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func fordemo2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}

func fordemo3() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func main() {
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
	// fordemo()
	// fordemo2()
	fordemo3()
}
