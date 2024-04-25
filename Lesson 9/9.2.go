package main

import "fmt"

//Необходимо создать срез целых чисел: 1, 2, 3. Далее создать 4
//вложенных цикла. Каждый цикл должен в консоль выводить текущее значение среза. Причём внутренние срезы должны содержать отступы для облегчения визуального восприятия. Внутренний срез на ключе 1 должен остановить все циклы, начиная со второго цикла. В консоли должно быть:
//
//v1: 1
//v2: 1
//v3: 1
//v4: 1
//v4: 2
//v1: 2
//v2: 1
//v3: 1
//v4: 1
//v4: 2
//v1: 3
//v2: 1
//v3: 1
//v4: 1
//v4: 2

func main() {
	sliceHomeWork := []int{1, 2, 3}

	for i := 0; i < 3; i++ {
		fmt.Println("v1:", sliceHomeWork[i])
	SOME_LABEL:
		for j := 0; j < 3; j++ {
			fmt.Println("\tv2:", sliceHomeWork[j])
			for k := 0; k < 3; k++ {
				fmt.Println("\t\tv3:", sliceHomeWork[k])
				for l := 0; l < 3; l++ {
					fmt.Println("\t\t\tv4:", sliceHomeWork[l])
					if l == 1 {
						break SOME_LABEL
					}
				}
			}
		}
	}
}
