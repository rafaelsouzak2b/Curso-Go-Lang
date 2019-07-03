package main

import "fmt"

func main(){
	i := 1
	//Go não tem aritmetica de ponteiro

	var p *int = nil

	p = &i // pega o endereco de i e atribui ao ponteiro

	*p++ //desreferenciando (pegando o valor)
	i++

	fmt.Println(p, *p, i, &i)
}