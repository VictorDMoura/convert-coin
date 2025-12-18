package main

import (
	"encoding/json"
	"fmt"
	"io"
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
	rate, err := getRate(coinType)
	if err != nil {
		fmt.Printf("Error fetching exchange rate: %v\n", err)
		return
	}
	convertedValue := value * rate

	fmt.Printf("Converted value: %.2f %s\n", convertedValue, coinType)
}



func getRate(coinType string) (float64, error) {
	url := fmt.Sprintf("https://economia.awesomeapi.com.br/json/last/%s-BRL", coinType)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == 429 {
			return 0, fmt.Errorf("erro e429: The api rate limit has been exceeded")
		}
		return 0, fmt.Errorf("The api returned status error: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var rateData map[string]ExchangeRates
	err = json.Unmarshal(body, &rateData)
	if err != nil {
		return 0, fmt.Errorf("error parsing JSON: %v", err)
	}

	keyValue := fmt.Sprintf("%sBRL", coinType)
	rate, exists := rateData[keyValue]
	if !exists {
		return 0, fmt.Errorf("rate for '%s' not found", coinType)
	}

	return strconv.ParseFloat(rate.Bid, 64)
}