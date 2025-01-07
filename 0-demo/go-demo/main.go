package main

import (
	"fmt"
	"math"
)

func main() {
	const bmiPower int = 2

	var userHeight float64 = 1.74
	var userWeight float64 = 93.9
	var userBmi float64 = userWeight / math.Pow(userHeight, float64(bmiPower))

	fmt.Printf("user's bmi is %.2f\n", userBmi)
}
