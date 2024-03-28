package main

import "fmt"

func main() {

	const (
		lala = iota ^ 100
		lolo
		test = 30
		test2
	)
	fmt.Print(lala, lolo, test, test2)
}
