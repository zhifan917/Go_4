package main

import "fmt"

var isWhiteUser = map[int]bool{
	11123121: true,
	12312121: true,
	11111111: false,
	73273737: true,
}

func isWhite(userID int) bool {
	_, ok := isWhiteUser[userID]
	return ok
}

func main() {
	userID := 11111111
	if isWhite(userID) {
		fmt.Println("user is ok")
	} else {
		fmt.Println("user is failed")
	}
}
