package main

import "fmt"

// func main() {
// 	s1 := "哈登niub"
// 	// s2 := 'G'
// 	// fmt.Println(s1, s2)
// 	for i := 0; i < len(s1); i++ {
// 		fmt.Printf("%v(%c)\n ", s1[i], s1[i])
// 	}

// 	for _, v := range s1 {
// 		fmt.Printf("%v(%c) ", v, v)

// 	}
// }

func main() {
	s1 := "big"
	bytes1 := []byte(s1)
	bytes1[0] = 'p'
	fmt.Println(string(bytes1))

	s2 := "张志凡"
	bytes2 := []rune(s2)
	bytes2[2] = '傻'
	fmt.Println(string(bytes2))
}
