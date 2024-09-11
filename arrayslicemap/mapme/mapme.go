package main

import "fmt"

func main() {
	aprovados := make(map[int]string)
	ACpf := make([]int, 0)
	ACpf = append(ACpf, 260307,122482,4965,8922,4002)
	AName := make([]string, 0)
	AName = append(AName, "Nicolas", "Gabriela", "Maria", "Ana", "Priscila")

	for i, cpf := range ACpf {
		aprovados[cpf] = AName[i]
	}
	for key, value := range aprovados{
		fmt.Printf("CPF: %d, Nome: %s\n", key, value)
	}
	fmt.Println(aprovados)
}