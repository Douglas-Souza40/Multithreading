package gateway

import (
	"encoding/json"
	"fmt"
	"github.com/Douglas-Souza40/Multithreading.git/internal/domain"
	"io"
	"net/http"
	"time"
)

func FetchFromBrasilAPI(cep string, ch chan<- domain.Address) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	client := http.Client{Timeout: 1 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var raw domain.BrasilAPIResponse
	if err := json.Unmarshal(body, &raw); err != nil {
		return
	}

	ch <- domain.Address{
		CEP:        raw.CEP,
		Logradouro: raw.Street,
		Bairro:     raw.Neighborhood,
		Localidade: raw.City,
		UF:         raw.State,
		Source:     "BrasilAPI",
	}
}

func FetchFromViaCEP(cep string, ch chan<- domain.Address) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	client := http.Client{Timeout: 1 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var raw domain.ViaCEPResponse
	if err := json.Unmarshal(body, &raw); err != nil {
		return
	}

	ch <- domain.Address{
		CEP:        raw.CEP,
		Logradouro: raw.Logradouro,
		Bairro:     raw.Bairro,
		Localidade: raw.Localidade,
		UF:         raw.UF,
		Source:     "ViaCEP",
	}
}
