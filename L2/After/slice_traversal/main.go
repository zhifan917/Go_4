package main

import "fmt"

func main() {
	s := []int{1, 3, 5}

	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	for k, v := range s {
		fmt.Println(k, v)
	}
}
