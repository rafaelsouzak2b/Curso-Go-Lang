package main
import (
	"time"
	"fmt"
)

func main(){
	t := time.Now()
	switch{ // switch true, pega o case que for verdadeiro
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}
}