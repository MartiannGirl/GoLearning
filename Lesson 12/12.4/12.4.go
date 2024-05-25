package main

import (
	"errors"
	"fmt"
)

//В представленном ниже коде возникает паника. Нужно понять где и объяснить почему.
//Также нужно предложить не менее 3-х вариантов решения проблемы с объяснением причины устранения ошибки

type Bird interface {
	Sing() string
}
type Duck struct {
	voice string
}

func (d Duck) Sing() string {
	return d.voice
}
func main() {
	d := &Duck{"Течет ручей, бежит ручей"}
	//var d *Duck единственное, что я понимаю - это неверный указатель на структуру
	song, err := Sing(d)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)
}
func Sing(b Bird) (string, error) {
	if b != nil {
		return b.Sing(), nil
	}
	return "", errors.New("Ошибка пения!")
}
