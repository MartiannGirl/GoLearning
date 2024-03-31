package main

import "fmt"

const (
	ImDeadInside = 1
	ImAlive      = 0
	Two          = 2
	Three        = 3
	Chetyre      = 4
)

const nextTask int = 1
const templateLetter string = "Приветики-пистолетики" // global

func main() {
	const IdUser int = 2 // local
	fmt.Println("This User is:", IdUser)
	fmt.Println("Приветствие: ", templateLetter)

	const nextTask int = 42
	fmt.Println("Ответ на главный вопрос человечества: ", nextTask) // task 3.3

	fmt.Println("Task 3.4: ", ImDeadInside, ImAlive, Two, Three, Chetyre) // task 3.4

	const (
		MuriloBenicio     = 0
		GiovannaAntonelli = 1
		VeraFischer       = 2
		ReginaldoFaria    = 3
		JandiraMartini    = 4
	)
	fmt.Println("Task 3.4: ", MuriloBenicio, GiovannaAntonelli, VeraFischer, ReginaldoFaria, JandiraMartini) // task 3.5

	const n int = 5
	var Param float32 = 3.4 + float32(n)

	fmt.Println(Param)

	const (
		a = iota + 1
		b = 1 << iota
		c
		d
		f
	)

	fmt.Println(a, b, c, d, f)

}
