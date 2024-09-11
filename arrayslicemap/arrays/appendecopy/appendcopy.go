package main

import "fmt"

func main() {
	srcSlice := []int{4, 5, 6, 7}
	destSlice := make([]int, 5)

	numCopied := copy(destSlice, srcSlice)
	fmt.Println("Número de elementos copiados:", numCopied)
	fmt.Println("Slice de origem:", srcSlice)
	fmt.Println("Slice de destino:", destSlice)
}
