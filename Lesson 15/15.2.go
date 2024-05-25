package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var j int32
	for i := 1; i < 15; i++ {
		go func() {
			atomic.AddInt32(&j, 1)

		}()
		fmt.Println(j)
	}

}
