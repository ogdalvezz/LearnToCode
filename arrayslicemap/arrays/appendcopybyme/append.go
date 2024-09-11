package main

import "fmt"

func main() {
	cpflist := make([]int, 5)
	cpflist[0], cpflist[1], cpflist[2], cpflist[3], cpflist[4] = 122486, 122482, 260307, 26262, 34912
	aprovado := make(map[int]int)
	for i, cpf := range cpflist {
		aprovado[i] = cpf
	}
	for key, value := range aprovado {
		fmt.Printf("Índice: %d, CPF: %d\n", key, value)
	}
	fmt.Println(aprovado," novos")
}
