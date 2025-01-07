package main

import "fmt"

func main() {
	var usdToEur float64 = 0.95
	var usdToRub float64 = 100
	eurToRub := usdToRub / usdToEur
	fmt.Printf("eurToRub: %.2f\n", eurToRub)
}
