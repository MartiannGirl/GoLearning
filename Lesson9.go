package main

import "fmt"

/*type kolvo int // недоделано 9.1


func (k kolvo) fruitMarket() string {
	m1 := map[string]int{}
	m1["апельсины"] = 5
	m1["яблоки"] = 3
	m1["сливы"] = 1
	m1["груши"] = 0

	return fruitMarket()
}
func fruitMarket() string {
	return "lalala"
}
func main() {

	fmt.Println("Количество: ", fruitMarket())
	fmt.Println(fruitMarket())
}

*/
/*
func main() {
	s1 := []int{1, 2, 3}
	//fmt.Println("s1[1]:", &s1[1])
	//fmt.Println("s1:", s1, "длина:", len(s1), "ёмкость:", cap(s1))
	sum := 0
	for i := 0; i < 10; i++ {
		sum++
		fmt.Println("v: ", i, s1)
	}
	//fmt.Println("Сумма: ", sum)
}
*/
func main() { // 9.2 еще не разобралась до конца
	s1 := []int{1, 2, 3}
	for i := 0; i < len(s1); i++ {
		//fmt.Println(s1[i], len(s1))
	SOME_LABEL:
		for j := 0; j < 3; j++ {
			//fmt.Println(s1[j])
			for n := 0; n < 2; n++ {
				//if n == 1 {
				//	break SOME_LABEL
				//}
				for k := 0; k < 1; k++ {
					fmt.Println("v", s1[j])
					if k == 4 {
						break SOME_LABEL
					}
				}
			}
		}
	}
}
