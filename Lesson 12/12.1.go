package main

import (
	"fmt"
	"log"
)

func main() {
	a := 1
	do(a)

}
func do(v any) {
	a := 10
	v = 10

	//a += int(v)

	/*
	   Здесь нужно увеличить значение a на v.
	   В случае невозможности приведения к int
	   необходимо сообщить об этом и немедленно
	   завершить полнение программы.
	*/

	w, ok := v.(int)
	if !ok {
		log.Fatalln("please stop it")
	}
	a += w
	fmt.Println(a)
}
