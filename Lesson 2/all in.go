package main

import (
	"fmt"
)

func main() {
	var a = 16
	var b = 3
	var c = 16 / 3
	var d = a % b
	//var d =  ('Результат: ', c, '', 'остаток от деления:', '', 'тип результата: '}

	fmt.Printf("Результат: %d", c)
	fmt.Printf(", остаток от деления: %d", d)
	fmt.Printf(", тип результата: %T", c)

}
