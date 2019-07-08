package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f\n", media(7.7, 8.1, 5.9, 9.9))
	fmt.Printf("Média: %.2f\n", media())
	fmt.Printf("Média: %.2f\n", media(10.20, 5.84, 2.46, 74.20, 20.41, 1.54))
	fmt.Printf("Média: %.2f\n", media(8.2, 10.0))
}
