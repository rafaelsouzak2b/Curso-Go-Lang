package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	//array1 = append(array1, 4, 5, 6) //não é possivel utilizar append no array
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	//copy não aumenta a capacidade do slice, apenas copia
	copy(slice2, slice1) //copia para o slice2 o conteudo do slice1
	fmt.Println(slice2)  // imprimi o que foi passado para o slice2, nesse caso copia apenas os dois primeiros valores, pois o slice2 suporte 2 posições

}
