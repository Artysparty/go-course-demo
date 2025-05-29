package main

import (
	"errors"
	"fmt"
	"math"
)

func checkIMT() {
	fmt.Println("__ Калькулятор массы тела __")

	for {
		userHeight, userWeigth := getUserInput()
		IMT, err := calculateIMT(userHeight, userWeigth)
		if err != nil {
			fmt.Println("Не заданы параметры расчета")
			continue
		}
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

func calculateIMT(height, weight float64) (float64, error) {
	if height <= 0 || weight <= 0 {
		return 0, errors.New("NO_PARAMS_ERROR")
	}
	const IMTPower = 2
	IMT := weight / math.Pow(height/100, IMTPower)
	return IMT, nil
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
