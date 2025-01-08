package main

import (
	"fmt"
	"math"
)

func main() {

	userWeight, userHeight := getUserInput()

	userBmi := calculateBmi(userWeight, userHeight)

	if userBmi < 16.0 {
		fmt.Println("bmi is too small")
	} else if userBmi >= 16.0 && userBmi < 18.5 {
		fmt.Println("body mass deficit")
	} else if userBmi >= 18.5 && userBmi < 25.0 {
		fmt.Println("bmi is okay")
	} else if userBmi >= 25.0 && userBmi < 30 {
		fmt.Println("body mass slightly exceeds target")
	} else {
		fmt.Println("body mass significantly exceeds target")
	}

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
