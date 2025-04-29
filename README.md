# 📦 Multithreading (Go)

Esta aplicação em Go busca o endereço de um CEP consultando **duas APIs simultaneamente**:
- [BrasilAPI](https://brasilapi.com.br/)
- [ViaCEP](https://viacep.com.br)

A resposta mais rápida entre elas é utilizada. Caso nenhuma responda em até 1 segundo, é exibido um erro de timeout.

---

## 🚀 Funcionalidades

- Busca paralela em duas APIs distintas.
- Considera apenas a **resposta mais rápida**.
- Timeout configurado para **1 segundo**.
- Exibe resultado formatado no terminal.
- Arquitetura baseada em **Clean Architecture**.
- Suporte a testes com TDD.

---

## 🧱 Estrutura do Projeto

Multithreading/ ├── cmd/ │ └── main.go # Ponto de entrada da aplicação ├── internal/ │ ├── domain/ # Models do domínio │ ├── usecase/ # Lógica de negócio (casos de uso) │ └── infra/ │ └── gateway/ # Comunicação com as APIs externas ├── test/ # Testes unitários ├── go.mod # Dependências Go └── README.md


---

## 🛠️ Como executar localmente

### 1. Clone o repositório

```bash
git clone https://github.com/seu-usuario/cep-fetcher.git
cd cep-fetcher
```

### 2. Inicialize o projeto Go
```
go mod tidy
```
### 3. Execute a aplicação com o CEP desejado
```
go run cmd/main.go 01153000
```
### 4. Exemplo de saída:
```
Resposta mais rápida (ViaCEP): {CEP:01153-000 Logradouro:Rua Exemplo Bairro:Centro Localidade:São Paulo UF:SP Source:ViaCEP}
```