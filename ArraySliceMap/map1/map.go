package main

import "fmt"

func main() {
	//var aprovados map[int]string
	//mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[12345678468] = "Pedro"
	aprovados[12349378978] = "Ana"

	fmt.Println(aprovados)
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12349378978])
	delete(aprovados, 12349378978)
	fmt.Println(aprovados[12349378978])
}
