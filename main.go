package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("__ Калькулятор массы тела __")

	for {
		userHeight, userWeigth := getUserInput()
		IMT := calculateIMT(userHeight, userWeigth)
		outputResult(IMT)
		isContinue := checkContinue()
		if !isContinue {
			return
		}
	}
}

func checkContinue() bool {
	var userAnswer string

	fmt.Print("Посчиать индекс массы тела еще раз? (y/n): ")
	fmt.Scan(&userAnswer)

	return userAnswer == "y" || userAnswer == "Y"
}

func outputResult(IMT float64) {
	switch {
	case IMT < 16:
		fmt.Println("У вас критический недостаток веса")
	case IMT < 18.5:
		fmt.Println("У вас недостаток веса")
	case IMT < 25:
		fmt.Println("У вас нормальный вес")
	case IMT < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас ожирение")
	}
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	fmt.Println(result)
}

func calculateIMT(height, weight float64) float64 {
	const IMTPower = 2
	IMT := weight / math.Pow(height/100, IMTPower)
	return IMT
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userWeigth float64

	fmt.Print("Введите свой рост в см: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userWeigth)

	return userHeight, userWeigth
}
