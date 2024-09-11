package main

import "fmt"

func main() {
	var msgenviada [4]int
	fmt.Println(msgenviada)
	msgenviada[0], msgenviada[1], msgenviada[2], msgenviada[3] = 1, 2, 3, 4
	fmt.Println(msgenviada)
	tokenmsg := len(msgenviada)
	fmt.Println("tokens restantes: ", tokenmsg)
	for i := 0; i < len(msgenviada); i++ {
		if tokenmsg > 0 {
			fmt.Println(msgenviada[i], "mensagem enviada")
			tokenmsg -= 1
			fmt.Println(tokenmsg, "tokens")
			fmt.Println("mais uma mensagem enviada com sucesso")
		} else if tokenmsg <= 0{
			fmt.Println("deu erro pois nao é maior que 0")
		}
	}
}
