package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	//------------------tipos numericos inteiros-----------------------------

	fmt.Println(1, 2, 1000)

	//reflect.TypeOf(32000) - retorna o tipo do valor passado
	fmt.Println("Literal inteiro é:", reflect.TypeOf(32000))

	//******************sem sinal {só possitivos}... uint8 uint16 uint32 uint64 ***************

	var b byte = 255 //tipo uint8
	fmt.Println("O byte é:", reflect.TypeOf(b))

	//**********com sinal.. int8 int16 int32 int64 ************

	//math.MaxInt64 - pega o maior valor inteiro possivel de 64 bits
	i1 := math.MaxInt64

	fmt.Println("O valor maximo de int é:", i1)
	fmt.Println("O tipo de i1 é:", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é:", reflect.TypeOf(i2))
	fmt.Println("O valor de i2 na tabela Unicode é:", i2)

	//------------------números reais {float32, float64}-----------------------------

	var x float32 = 49.99 // outro forma var x = float32(49.99)
	fmt.Println("O tipo de x é:", reflect.TypeOf(x))
	fmt.Println("O tipo literal de 49.99 é:", reflect.TypeOf(49.99)) //por padrão os valores são em float64

	//------------------tipo boolean (não aceita inteiro)-----------------------------
	bo := true
	fmt.Println("O tipo de bo é:", reflect.TypeOf(bo))
	fmt.Println(!bo) // imprime a negação de bo

	//------------------tipo string-----------------------------
	s1 := "Ola meu nome é Rafael"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é:", len(s1))

	//strings com multiplas linhas
	s2 := `Ola
	meu
	nome
	é
	Rafael
	`
	fmt.Println("O tamanho da string é:", len(s2))

	//------------------tipo char (char é um int32)-----------------------------

	char := 'a'
	fmt.Println("O tipo de char é:", reflect.TypeOf(char))
	fmt.Println("O valor de char é:", char)

}
