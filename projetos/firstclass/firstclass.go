package main

import "fmt"

func executarOperacao(x, y int, operacao func(int, int) int) int {
	return operacao(x, y)
}
func main() {
	soma := func(a, b int) int {
		return a + b
	}
	resultado := executarOperacao(7, 8, soma)
	fmt.Println(resultado)
}
