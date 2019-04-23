package main

import "fmt"

func main() {
	s1 := make([]int, 3)
	s2 := s1
	s2[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)
}
