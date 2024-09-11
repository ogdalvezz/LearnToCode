package main

import "fmt"

func main() {
	//var aprovados map[int]string
	//mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[1] = "Maria"
	aprovados[2] = "ju"
	aprovados[3] = "lia"
	aprovados[4] = "ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados{
		fmt.Printf("%s (CPF:%d)\n", nome, cpf)
	}

	fmt.Println(aprovados[1])
	delete(aprovados, 1)
	fmt.Println(aprovados)
}