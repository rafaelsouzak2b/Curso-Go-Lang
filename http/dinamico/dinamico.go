package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05") //Formatação: d = 02, m = 01, Y = 2006, h = 03, m = 04, s = 05
	fmt.Fprintf(w, "<h1>Hora certa: %s</h1>", s)  //escreve no w o trecho html
}

func main() {
	http.HandleFunc("/horacerta", horaCerta)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
