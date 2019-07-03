package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 // inferencia de tipo
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 2 // i = i / 2
	i *= 2 // i = i * 2
	i %= 2 // i = i % 2

	fmt.Println(i)

	x, y := 1, 2 // inicializa dois variaveis e atribui valores a elas
	fmt.Println(x, y)

	x, y = y, x // inverte os valores das variaveis
	fmt.Println(x, y)
}
