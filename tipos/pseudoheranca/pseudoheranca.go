package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       // campos anonimos, passa apenas o tipo
	turboLigado bool
}

func main() {
	f := ferrari{}
	//Com o campo anonimo você consegue acessar diretamente os campos de outro struct
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Printf("A ferrari %s está com turbo ligado? %v\n", f.nome, f.turboLigado)
	fmt.Println(f)
}
