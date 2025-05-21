package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2

	var userHeight float64
	var userWeigth float64

	fmt.Println("__ Калькулятор массы тела __")
	fmt.Print("Введите свой рост в см: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userWeigth)

	IMT := userWeigth / math.Pow(userHeight/100, IMTPower)

	fmt.Printf("Ваш индекс массы тела: %.0f", IMT)
}
