package usecase

import (
	"errors"
	"fmt"
	"github.com/Douglas-Souza40/Multithreading.git/internal/domain"
	"github.com/Douglas-Souza40/Multithreading.git/internal/infra/gateway"
	"time"
)

func FetchFastestCEP(cep string) error {
	resultChan := make(chan domain.Address, 2)

	go gateway.FetchFromBrasilAPI(cep, resultChan)
	go gateway.FetchFromViaCEP(cep, resultChan)

	select {
	case result := <-resultChan:
		fmt.Printf("Resposta mais rápida (%s): %+v\n", result.Source, result)
		return nil
	case <-time.After(1 * time.Second):
		return errors.New("timeout: nenhuma API respondeu em 1 segundo")
	}
}
