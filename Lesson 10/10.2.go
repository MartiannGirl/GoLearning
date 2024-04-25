// Написать программу, которая импортирует пакет first задачи 8.1, но не использует его логику

// package second
package main

import "fmt"

func main() {
	fmt.Println("вот такие пироги")
}

// chmod +x my_second.go
// mkdir /usr/local/go/src/second
// cp my_second.go /usr/local/go/src/second/my_second.go
// go install second13
// sudo chmod -R 777 /usr/local/go/src/second
// ls /usr/local/go/src
