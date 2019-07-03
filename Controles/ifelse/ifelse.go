package main

import "fmt"

func imprimirResultado(nota float64){
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	}else{
		fmt.Println("Reprovado com nota", nota)
	}
}

func main(){
	imprimirResultado(8.9)
	imprimirResultado(4.8)
}