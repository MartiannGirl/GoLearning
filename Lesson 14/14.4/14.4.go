package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	stop := make(chan struct{}, 1)

	go func() {
		<-ch
		stop <- struct{}{}

	}()
	//<-stop

	fmt.Println("happy end")

}

/*
Не меняя логику дочерней горутины, нужно дописать6 логику главной горутины так,
чтобы ошибки блокировки не возникало.
Вместо ошибки в консоль должна выводится фраза «happy end».
Так как канал ch является «особым» , то в него ЗАПРЕЩАЕТСЯ писать значения.

возникнет ошибка блокировки:
fatal error: all goroutines are asleep - deadlock!
goroutine 1 [chan receive]:
main.main()
/go/src/course/main.go:12 +0x9c
goroutine 18 [chan receive]:
main.main.func1()
/go/src/course/main.go:9 +0x2c
created by main.main in goroutine 1
/go/src/course/main.go:8 +0x90
exit status 2
*/
