// /*
// const pi = 3.1415
// const e = 2.7182
// */

// // const (
// // 	const pi = 3.1415
// //     const e = 2.7182
// // )

// package main

// import "fmt"

// // const (
// // 	n1 = 100
// // 	n2
// // 	n3
// // )

// // func main() {
// // 	fmt.Println(n1, n2, n3)
// // }

// // const (
// // 	n1 = 3
// // 	n2 = iota //1
// // 	n3        //2
// // 	_
// // 	n4 //4
// // )

// // const n5 = iota //0

// const (
// 	_  = iota //0
// 	KB = 1 << (10 * iota) //1 << 10
// 	MB = 1 << (10 * iota) //1 << 20
// 	GB = 1 << (10 * iota) //1 << 30
// 	TB = 1 << (10 * iota) //1 << 40
// 	PB = 1 << (10 * iota) //1 << 50
// )

// func main() {
// 	fmt.Println(KB, MB, GB, TB, PB)
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// const (
// 	a = math.MaxFloat32
// 	b = math.MaxFloat64
// )

// func main() {
// 	fmt.Println(a, b)
// }

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.3f\n", math.Pi)
}
