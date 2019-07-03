package main

import "fmt"

func main() {

	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println("Nova")
	fmt.Println("linha.")

	x := 3.141516

	xs := fmt.Sprint(x) //retorna a string

	fmt.Println("O valor de x é: " + xs) // aenas concatena strings
	fmt.Println("O valor de x é:", x)    // pode concatenar sem precisar converter

	fmt.Printf("O valor de x é: %.2f.", x) //print formatado

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%d %.1f %t %s", a, b, c, d) // d para int, f para float, t para bool, s para string

	fmt.Printf("\n%v %v %v %v", a, b, c, d) // v formato mais generico para varios tipos de valores

}
