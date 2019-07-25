package main

import (
	"fmt"

	"github.com/rafaelsouzak2b/area"
)

func main() {
	fmt.Println(area.Circ(6.0))
	fmt.Println(area.Rect(5.0, 2.0))
	//fmt.Println(area._TrianguloEq(5.0, 2.0)) Não funciona pois _TrianguloEq não é público
}
