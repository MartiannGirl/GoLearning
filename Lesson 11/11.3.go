package main

import (
	"errors"
	"fmt"
)

//«ошибка3:ошибка2:ошибка1». Не используя unwrap, нужно определить была ли ошибка «ошибка1»

func main() {
	//err := "Ошибка4"

	a := errors.New("ошибка1:")
	b := fmt.Errorf("ошибка2:%w", a)
	//c := fmt.Errorf("ошибка1:%w", b)

	//fmt.Println(a)

	fmt.Println("ошибка1:", errors.Is(b, a))

	//d := fmt.Errorf(err)
	//fmt.Println(d)

}
