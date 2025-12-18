# Convert Coin CLI 💰

![Go Version](https://img.shields.io/badge/go-1.25.4-blue) ![License](https://img.shields.io/badge/license-MIT-green)

[🇺🇸 English](#english) | [🇧🇷 Português](#português)

---

<a name="english"></a>
## 🇺🇸 English

### Description
**Convert Coin** is a Command Line Interface (CLI) tool built in Go. It fetches real-time exchange rates using the AwesomeAPI to convert currency values accurately and instantly.

### Features
- 🚀 **Real-time Data:** Fetches the latest exchange rates via HTTP requests.
- 🛡️ **Robust:** Handles API rate limits (429 errors) and connection issues.
- ⚡ **Fast:** Lightweight executable using standard Go libraries.

### Prerequisites
- [Go](https://go.dev/dl/) installed on your machine.

### Installation & Usage

1. **Clone the repository:**
   ```bash
   git clone [https://github.com/VictorDMoura/convert-coin.git](https://github.com/VictorDMoura/convert-coin.git)
   cd convert-coin
    ```

2. **Run directly:**

Pass the amount and the currency code you want to check against BRL (Brazilian Real).
```bash
go run main.go <amount> <currency>

```


3. **Build and Run (Optional):**
```bash
go build -o convert
./convert 10 USD

```



### Example

Converting 10 units of USD (United States Dollar):

```bash
$ go run main.go 10 USD
Converted value: 51.50 USD

```

---

<a name="português"></a>

## 🇧🇷 Português

### Descrição

**Convert Coin** é uma ferramenta de linha de comando (CLI) desenvolvida em Go. Ela busca taxas de câmbio em tempo real utilizando a AwesomeAPI para converter valores monetários de forma precisa e instantânea.

### Funcionalidades

* 🚀 **Dados em Tempo Real:** Busca as taxas de câmbio mais recentes via requisições HTTP.
* 🛡️ **Robusto:** Lida com limites de requisição da API (erros 429) e problemas de conexão.
* ⚡ **Rápido:** Executável leve usando apenas as bibliotecas padrão do Go.

### Pré-requisitos

* [Go](https://go.dev/dl/) instalado em sua máquina.

### Instalação e Uso

1. **Clone o repositório:**
```bash
git clone [https://github.com/VictorDMoura/convert-coin.git](https://github.com/VictorDMoura/convert-coin.git)
cd convert-coin

```


2. **Rodando diretamente:**
Passe o valor e o código da moeda que deseja consultar em relação ao BRL (Real Brasileiro).
```bash
go run main.go <valor> <moeda>

```


3. **Compilar e Rodar (Opcional):**
```bash
go build -o convert
./convert 10 USD

```



### Exemplo

Convertendo 10 unidades de USD (Dólar Americano):

```bash
$ go run main.go 10 USD
Converted value: 51.50 USD

```



