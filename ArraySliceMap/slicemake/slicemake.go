package main

import "fmt"

func main() {
	s := make([]int, 10) // criar um slice sem referenciar um array, ele cria um array interno
	s[9] = 12

	fmt.Println(s)

	s = make([]int, 10, 20) // cria um slice de 10 posições e um array interno de 20 posições

	fmt.Println(s, len(s), cap(s)) //cap() pega a capacidade máxima do slice

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)

	fmt.Println(s, len(s), cap(s))

	s = append(s, 1) // se  tamanho do slice passar o tamanho máximo do array interno, esse array aumenta de tamanho
	fmt.Println(s, len(s), cap(s))

}
