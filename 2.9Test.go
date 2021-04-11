package main

import (
	"fmt"
	"math/rand"
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
		pausota()
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
		pausota()
		//pausita()
	}
}

func main() {
	go p()
	go q()

	time.Sleep(3 * time.Second)
	fmt.Println(n)
}

func pausita() {
	d := rand.Intn(50) + 50
	time.Sleep(time.Nanosecond * time.Duration(d))
}
func pausota() {
	time.Sleep(time.Second * 4)
}
