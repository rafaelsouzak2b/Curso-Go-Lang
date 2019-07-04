package main

import (
	"fmt"
)

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} //o compilador conta o tamanho do array

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	for _, num := range numeros { // não utilizando o índice
		fmt.Println(num)
	}

}
