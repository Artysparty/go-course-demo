package main

import (
	"fmt"
)

func arraySlice() {
	fmt.Println("Бюджетный калькулятор")

	transactions := scanTransactions()
	balance := calculateTransactions(transactions)

	fmt.Printf("Итоговый баланс: %v \n", balance)
}

func scanTransactions() []float64 {
	transactions := []float64{}

	for {
		var transaction float64
		fmt.Print("Введите транзакцию (n для выхода): ")
		fmt.Scan(&transaction)

		if transaction == 0 {
			break
		}

		transactions = append(transactions, transaction)
	}

	return transactions
}

func calculateTransactions(transactions []float64) float64 {
	var balance float64

	for _, value := range transactions {
		balance = balance + value
	}

	return balance
}
