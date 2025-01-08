package main

import "fmt"

func main() {
	value := getUserInput()
	fmt.Println(calculateOutput(value, "EUR", "RUB"))
}

func getUserInput() float64 {
	var value float64
	fmt.Println("Enter a number: ")
	_, err := fmt.Scan(&value)
	if err != nil {
		fmt.Println(err)
	}
	return value
}

func calculateOutput(value float64, sourceCurrency string, targetCurrency string) string {
	var usdToEur float64 = 0.95
	var usdToRub float64 = 100
	eurToRub := usdToRub / usdToEur
	fmt.Printf("from %s to %s\n", sourceCurrency, targetCurrency)
	fmt.Printf("eurToRub: %.2f\n", eurToRub)
	return fmt.Sprintf("%.2f", value*eurToRub)
}
