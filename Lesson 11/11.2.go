package main

import (
	"errors"
	"fmt"
)

//«ошибка3:ошибка2:ошибка1». Из созданной цепочки ошибок нужно получить ошибку «ошибка2:ошибка1» и вывести в stdout

func main() {

	a := errors.New("Ошибка1:")
	a = fmt.Errorf("Ошибка2:%w", a)
	a = fmt.Errorf("Ошибка3:%w", a)
	//fmt.Println(a)

	d := errors.Unwrap(a)
	fmt.Println(d)

}
