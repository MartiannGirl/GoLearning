package main

import (
	"errors"
	"fmt"
)

// Нужно создать, используя оборачивание, ошибку «ошибка3:ошибка2:ошибка1»
func main() {
	err := "вот такие пироги"

	a := errors.New("Ошибка3:")
	a = fmt.Errorf("Ошибка2:%w", a)
	a = fmt.Errorf("Ошибка1:%w", a)
	fmt.Println(a)

	d := fmt.Errorf(err)
	fmt.Println(d)

}
