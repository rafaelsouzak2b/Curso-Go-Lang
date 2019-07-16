package main

import (
	"fmt"
	"strconv"
)
func main(){
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	fmt.Println("Teste " + string(97)) // traz o caractere corespondente da tabela asc2

	//int para string 
	fmt.Println("Teste " + strconv.Itoa(123))

	//A função strconv retorna dois valores(a conversão e o erro)
	//string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true") // string para bool

	if b {
		fmt.Println("Verdadeiro")
	}

}