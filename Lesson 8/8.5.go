package main

import "fmt"

//Необходимо создать карту с элементами: «слон» = 3, «бегемот» = 0, «носорог» = 5, «лев» = 1.
//Изменить «бегемот» на 2 и вывести результат в консоль

func main() {

	map1 := map[string]int{}

	map1["слон"] = 3
	map1["бегемот"] = 0
	map1["носорог"] = 5
	map1["лев"] = 1

	fmt.Println("В Кении можно встретить таких животных, как: ", map1)
	map1["бегемот"] = 2
	fmt.Println("На Марсе можно встретить таких животных, как: ", map1)
}
