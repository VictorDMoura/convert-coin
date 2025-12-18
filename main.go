package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type ExchangeRates struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}



func main(){

	data, err := loadRates("./rates.json")
    if err != nil {
        fmt.Println("Erro fatal:", err)
        return
    }
	
	arguments := os.Args[1:]
	if len(arguments) != 2 {
		fmt.Println("The program must be run with exactly two arguments.")
		return
	}
	value, err := strconv.ParseFloat(arguments[0], 64)
	if err != nil {
		fmt.Println("Error: The first argument must be a valid float.")
		return
	}
	coinType := arguments[1]
	fmt.Printf("Value: %.2f, Coin Type: %s\n", value, coinType)
	fmt.Println("Taxa do USD carregada:", data.Rates["USD"])
}

func loadRates(filepath string) (ExchangeRates, error) {
	var rates ExchangeRates

	content, err := os.ReadFile(filepath)
	if err != nil {
		return rates, fmt.Errorf("error reading file: %v", err)
	}

	err = json.Unmarshal(content, &rates)
	if err != nil {
		return rates, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return rates, nil
}


