package main

import (
	"fmt"
	"os"
	"strconv"
)


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
	coinType := arguments[1]
	fmt.Printf("Value: %.2f, Coin Type: %s\n", value, coinType)
}