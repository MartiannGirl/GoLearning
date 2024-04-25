package main

import "fmt"

//Необходимо написать функцию fruitMarket, которая будет принимать название фруктов, а возвращать их количество.
//Например, «апельсины» - 5. Сама функция должна создать карту: апельсин=5, яблоки=3, сливы=1, груши=0.
//Если запрашиваемых фруктов нет в карте, то в консоль должно выводится сообщение об отсутствии

func fruitMarket(fm string) {

	market := map[string]int{}
	market["апельсин"] = 5
	market["яблоки"] = 3
	market["сливы"] = 1
	market["груши"] = 0

	val, ok := market[fm]
	if ok {
		fmt.Println(fm, val, "шт")
	} else {
		fmt.Println("Такого фрукта нет в наличии")

	}
}
func main() {
	fruitMarket("")
	fruitMarket("сливы")
}
