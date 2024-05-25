// Необходимо создать функцию start, которая в консоль выводит
// некоторое сообщение. Необходимо запустить 10 горутин, которые будут запускать функцию start
// и выводить в консоль факт своего запуска,
// причём необходимо обеспечить однократный запуск функции start
package main

import (
	"fmt"
	"sync"
	"time"
)

func start() string {
	return "Lets start!"

}

//	func main() {
//		b := start()
//		fmt.Println("Test", b)
//	}
func main() {
	var n int
	b := start()
	var mu1 sync.Mutex

	for i := 0; i < 10; i++ {
		go func() {
			mu1.Lock()

			defer mu1.Unlock()
			n++
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println(b, n)
}
