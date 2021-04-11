package main

import (
	"fmt"
	"time"
)

var n int

func p() {
	var temp int
	for i := 0; i < 10; i++ {
		temp = n
		//fmt.Println("P: temp=", temp)

		n = temp + 1
		//fmt.Println("P: n=", n)
		pausa()
		//pausita()
	}
}
func q() {
	var temp int
	for i := 0; i < 10; i++ {
		temp = n
		//fmt.Println("Q: temp=", temp)

		n = temp + 1
		//fmt.Println("Q: n=", n)
		pausa()
		//pausita()
	}
}

func main() {
	go p()
	go q()

	time.Sleep(3 * time.Second)
	fmt.Println(n)
}
func pausa() {
	time.Sleep(time.Second * 4)
}
