package main

import (
	"fmt"
	"github.com/Douglas-Souza40/Multithreading.git/internal/usecase"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <CEP>")
		return
	}
	cep := os.Args[1]

	if err := usecase.FetchFastestCEP(cep); err != nil {
		fmt.Println("Erro:", err)
	}
}
