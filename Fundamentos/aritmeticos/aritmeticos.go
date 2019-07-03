package main

import (
	"fmt"
	"math"
)

func main(){
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Modulo =>", a%b)

	// bitwise
	fmt.Println("E =>", a&b) // 11 & 10 = 10
	fmt.Println("Ou =>", a|b) // 10 | 11 = 11
	fmt.Println("Xor =>", a^b) // 10 ^ 11 = 01

	c := 2.0
	d := 3.0

	//outras operações utilizando math
	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Maior =>", math.Min(float64(c), float64(d)))
	fmt.Println("Maior =>", math.Pow(c, d))
	




}