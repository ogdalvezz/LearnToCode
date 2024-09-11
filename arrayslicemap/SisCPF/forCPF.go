package main

import "fmt"

func main() {
	listaPessoas := make(map[int]string)
	CPF := make([]int, 0)
	CPF = append(CPF, 260307,122482,4965,8922,4002)
	Nomes := make([]string, 0)
	Nomes = append(Nomes, "Nicolas", "Gabriela", "Maria", "Ana", "Priscila")

	for i, cpf := range CPF {
		listaPessoas[cpf] = Nomes[i]
	}
	for key, value := range listaPessoas{
		fmt.Printf("CPF: %d, Nome: %s\n", key, value)
	}
	fmt.Println(listaPessoas)
}