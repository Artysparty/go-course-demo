package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("__ Калькулятор массы тела __")
	userHeight, userWeigth := getUserInput()
	IMT := calculateIMT(userHeight, userWeigth)
	outputResult(IMT)
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	fmt.Print(result)
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
