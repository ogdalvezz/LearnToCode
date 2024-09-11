package main

import (
	"fmt"
)

func calcular(x, y int, operacao func(int, int) int) int {
	return operacao(x, y)
}
func main() {
	var i string
	var num1, num2 int
	somar := func(a, b int) int {
		return a + b
	}
	subtrair := func(a, b int) int {
		return a - b
	}
	fmt.Print("Você deseja subtrair ou somar: ")
	fmt.Scanln(&i)
	if i == "subtrair" {
		fmt.Print("Mande seu primeiro numero: ")
		fmt.Scanln(&num1)
		fmt.Print("Mande seu segundo numero: ")
		fmt.Scanln(&num2)
		resultado := calcular(num1, num2, subtrair)
		fmt.Println(resultado)
	} else if i == "somar" {
		fmt.Print("Mande seu primeiro numero: ")
		fmt.Scanln(&num1)
		fmt.Print("Mande seu segundo numero: ")
		fmt.Scanln(&num2)
		resultado := calcular(num1, num2, somar)
		fmt.Println(resultado)
	} else {
		fmt.Println("Nao existe essa funçao, execute a calculadora novamente!")
	}
}
