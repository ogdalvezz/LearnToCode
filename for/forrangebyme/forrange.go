package main

import "fmt"

func main() {
	notas := [...]float64{7.8, 4.3, 9.1}
	total := 0.0
	for _, nota := range notas {
		total += nota
		fmt.Println(total)
	}
}
