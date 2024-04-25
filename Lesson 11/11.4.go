package main

import (
	"errors"
	"fmt"
)

//Нужно создать, используя оборачивание, ошибку
//«ошибка3:ошибка2:ошибка1». Также нужно создать свою ошибку в виде структуры myFirstError,
//которая обязательно должна иметь метод Error() string.
//Необходимо убедиться, что в созданной цепочке ошибок не было ошибки типа myFirstError

//func main() {

//a = fmt.Errorf("Ошибка1:%w", a)
//a := errors.New("Ошибка3:")
//a = fmt.Errorf("Ошибка2:%w", a)
//fmt.Println(a)

type myFirstError string

func (e myFirstError) Error() string {
	return "Lalalala"
}
func main() {

	SomeErr1 := fmt.Errorf("Ошибка3:")
	SomeErr2 := fmt.Errorf("Ошибка2:%w", SomeErr1)
	SomeErr3 := fmt.Errorf("Ошибка1:%w", SomeErr2)

	//err1 := fmt.Errorf("некотораяошибка: %w", SomeErr3)

	err2 := myFirstError("Lalalala")
	//fmt.Println(SomeErr3, err2)
	fmt.Println("myFirstError:", errors.Is(SomeErr3, err2))

}
