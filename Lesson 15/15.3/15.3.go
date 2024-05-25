package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var k int
	var mu sync.Mutex
	for i := 0; i < 9; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			k++
		}()
	}
	time.Sleep(2 * time.Second)
	mu.Lock()
	defer mu.Unlock()
	fmt.Println(k)
}
