package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type ExchangeRates struct {
	Code string             `json:"code"`
	CodeIn string             `json:"codein"`
	Name string             `json:"name"`
	High string             `json:"high"`
	Low string             `json:"low"`
	VarBid string          `json:"varBid"`
	PctChange string       `json:"pctChange"`
	Bid string             `json:"bid"`
	Ask string             `json:"ask"`
	Timestamp string       `json:"timestamp"`
	CreateDate string      `json:"create_date"`
}



func main(){
	
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

	coinType := strings.ToUpper(arguments[1])
	rate, err := getRate(ExchangeRates{}, coinType)
	if err != nil {
		fmt.Printf("Error fetching exchange rate: %v\n", err)
		return
	}
	convertedValue := value * rate

	fmt.Printf("Converted value: %.2f %s\n", convertedValue, coinType)
}



func getRate(data ExchangeRates, coinType string) (float64, error) {

	var url string = fmt.Sprintf("https://economia.awesomeapi.com.br/json/last/%s-BRL", coinType)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var rateData map[string]ExchangeRates
	err = json.Unmarshal(body, &rateData)

	if err != nil {
		return 0, err
	}

	keyValue := fmt.Sprintf("%sBRL", coinType)

	rate, exists := rateData[keyValue]
	if !exists {
		return 0, fmt.Errorf("currency type '%s' not found", coinType)
	}

	return strconv.ParseFloat(rate.Bid, 64)
}