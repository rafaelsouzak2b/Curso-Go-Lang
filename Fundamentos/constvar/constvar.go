package main

import (
	"fmt"
	m "math" //pod colocar um alias para utilizar a biblioteca
)

func main() {

	const pi float64 = 3.1415 // cria uma constante
	var raio = 3.2            // cria uma variavel

	//forma reduzida de criar uma variavel
	area := pi * m.Pow(raio, 2)
	fmt.Println("A area da circunferencia é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)
	var e, f bool = true, false // declarar mais de uma variavel
	fmt.Println(e , f)
	g, h, i := 2, false, "epa" // declarar mais de uma variavel na forma reduzida
	fmt.Println(g, h, i)
}
