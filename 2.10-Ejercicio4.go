package main

import (
	"fmt"
)

var n int
var k int = 3

func p() {
	var temp int
	for i := 0; i < k; i++ {
		temp = n
		n = temp + 1
	}
}
func q() {
	var temp int
	for i := 0; i < k; i++ {
		temp = n
		n = temp - 1
	}
}

func main() {
	go p()
	go q()
	fmt.Println(n)
}
