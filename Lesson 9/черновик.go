package main

import (
	"fmt"
)

func main() {
	num := 5
	switch num {
	case 5:
		fmt.Println("пять")
	case 10:
		fmt.Println("меньшеилиравнодесяти")
	case 20:
		fmt.Println("двадцать")
	default:
		fmt.Println("значение:", num)
	}
}
