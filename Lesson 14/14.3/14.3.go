// Необходимо создать канал с буфером 4, записать в него 2 значения:
// «Привет», «буферизованный канал!». Далее необходимо прочитать все значения из канала и вывести в stdout
package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	var a string = "Привет"
	var b string = "буферизованный канал!"
	ch <- a
	ch <- b

	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
