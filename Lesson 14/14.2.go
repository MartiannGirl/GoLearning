package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		var n int
		n++
		fmt.Println("Привет, строковый канал!", n)
		ch <- n
	}()
	fmt.Println("Привет, строковый канал!", <-ch)
}
