package main

import (
	"fmt"
)

func main(){
	numeros := [...]int{1, 2, 3, 4, 5} //o compilador conta o tamanho do array

	for i, numeros := range numeros{
		fmt.Printf("%d) %d\n", i + 1, numeros)
	}
}