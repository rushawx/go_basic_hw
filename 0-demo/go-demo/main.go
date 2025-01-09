package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("__ body mass index __")

	for {
		userWeight, userHeight := getUserInput()

		userBmi, err := calculateBmi(userWeight, userHeight)
		if err != nil {
			panic(err)
		}

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

		isRepeat := checkRepeatCalculation()

		if !isRepeat {
			break
		}
	}
}

func outputResult(bmi float64) string {

	return fmt.Sprintf("user's bmi is %.2f", bmi)
}

func calculateBmi(userWeight float64, userHeight float64) (float64, error) {
	const bmiPower int = 2

	if userWeight <= 0 || userHeight <= 0 {
		return 0.0, errors.New("invalid input")
	}

	return userWeight / math.Pow(userHeight, float64(bmiPower)), nil
}

func getUserInput() (float64, float64) {
	var userWeight float64
	var userHeight float64

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

func checkRepeatCalculation() bool {
	var userChoice string
	fmt.Println("do you would like to repeat calculation?")
	_, err := fmt.Scan(&userChoice)
	if err != nil {
		fmt.Println(err)
	}
	if userChoice == "y" || userChoice == "Y" {
		return true
	}
	return false
}
