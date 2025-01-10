package main

import (
	"fmt"
	"log"
	"slices"
)

func main() {
	currencyCodes := []string{"EUR", "USD", "RUB"}
	currencyBank := map[string]float64{
		"EUR/USD": 1.05,
		"USD/EUR": 1 / 1.05,
		"RUB/EUR": 110.0,
		"EUR/RUB": 1 / 110.0,
		"RUB/USD": 100.0,
		"USD/RUB": 1 / 100.0,
	}
	fmt.Println(calculateOutput(currencyCodes, currencyBank))
}

func getCurrency(currencyCodes []string) (string, error) {
	fmt.Println("enter one of the following:")
	fmt.Println(currencyCodes)
	var currency string
	_, err := fmt.Scanln(&currency)
	if err != nil {
		log.Fatal(err)
	}
	if !slices.Contains(currencyCodes, currency) {
		return "", fmt.Errorf("currency %s not found in %v", currency, currencyCodes)
	}
	return currency, nil
}

func getValue() (float64, error) {
	var value float64
	fmt.Println("Enter a number: ")
	_, err := fmt.Scan(&value)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func calculateOutput(currencyCodes []string, currencyBank map[string]float64) string {
	fmt.Println("enter source currency code")
	sourceCurrency, err := getCurrency(currencyCodes)
	for err != nil {
		fmt.Println("Please enter a valid currency code: ")
		sourceCurrency, err = getCurrency(currencyCodes)
	}

	value, err := getValue()
	for err != nil {
		fmt.Println("Please enter a valid value: ")
		value, err = getValue()
	}

	fmt.Println("enter target currency code")
	targetCurrency, err := getCurrency(currencyCodes)
	for err != nil {
		fmt.Println("Please enter a valid currency code: ")
		targetCurrency, err = getCurrency(currencyCodes)
	}

	courseRate := currencyBank[targetCurrency+"/"+sourceCurrency]

	fmt.Printf("from %s to %s\n", sourceCurrency, targetCurrency)
	fmt.Printf("course rate: %.2f\n", courseRate)
	return fmt.Sprintf("%.2f", value*courseRate)
}
