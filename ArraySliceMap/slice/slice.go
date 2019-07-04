package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	//Slice não é um array! Slice define um pedaço do array.
	s2 := a2[1:3] // não cria um novo array, é uma estrutura que tem um ponteiro que aponta para o primeiro elemento pelo qual ele aponta

	fmt.Println(a2, s2)

	s3 := a2[:2] // novo slice, mas internamente ele esta apontando para o mesmo array da slice s2
	fmt.Println(a2, s3)

	//você pode imaginar um slice como: tamanho e um ponteiro para um elemento de um array

	s4 := s2[:1] //slice que um slice

	fmt.Println(s2, s4)

}
