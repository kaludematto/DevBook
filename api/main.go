package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando API!")

	r := router.Gerar()
	// chama a função Gerar criada em router e cria a rota usando http em localhost:5000
	log.Fatal(http.ListenAndServe(":5000", r))
}
