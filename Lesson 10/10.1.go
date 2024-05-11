package main

import (
	"first"
	"fmt"
)

func main() {
	msg := fmt.Sprintf("Hello, Dolly! ", first.Hello())

	fmt.Println(msg)
}
