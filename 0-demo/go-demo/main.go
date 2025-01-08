package main

import (
	"fmt"
	"math"
)

func main() {

	userWeight, userHeight := getUserInput()

	userBmi := calculateBmi(userWeight, userHeight)

	fmt.Println(outputResult(userBmi))
}

func outputResult(bmi float64) string {
	return fmt.Sprintf("user's bmi is %.2f", bmi)
}

func calculateBmi(userWeight float64, userHeight float64) float64 {
	const bmiPower int = 2

	return userWeight / math.Pow(userHeight, float64(bmiPower))
}

func getUserInput() (float64, float64) {
	var userWeight float64
	var userHeight float64

	fmt.Println("__ body mass index __")

	fmt.Println("enter your weight:")
	_, err := fmt.Scan(&userWeight)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("enter your height:")
	_, err = fmt.Scan(&userHeight)
	if err != nil {
		fmt.Println(err)
	}

	return userWeight, userHeight
}
