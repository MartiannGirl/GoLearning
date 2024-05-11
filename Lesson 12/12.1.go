package main

import (
	"log"
)

func main() {
	a := 1
	do(a)

}
func do(v any) {
	a := 10
	v = 10

	//c := a * v
	/*
	   Здесь нужно увеличить значение a на v.
	   В случае невозможности приведения к int
	   необходимо сообщить об этом и немедленно
	   завершить полнение программы.
	*/
	if a == 10 {

		log.Fatalf("Please stop it immediately", a)
	}
	//fmt.Println(a)
}
