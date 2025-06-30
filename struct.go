package main

import (
	"demo/app-1/account"
	"fmt"

	"github.com/fatih/color"
)

func structures() {
	fmt.Println("Менеджер паролей")
	vault := account.NewVault()

Menu:
	for {
		variant := getMenu()

		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		case 4:
			break Menu
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант: ")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Scan(&variant)
	return variant
}

func findAccount(vault *account.Vault) {
	url := scanTemplate("Введите url для поиска")

	accounts := vault.FindAccountByUrl(url)

	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}

	for _, account := range accounts {
		account.PrintAccount()
	}
}

func deleteAccount(vault *account.Vault) {
	url := scanTemplate("Введите url для удаления")

	isDeleted := vault.DeleteAccountByUrl(url)

	if isDeleted {
		color.Green("Удалено")
	} else {
		color.Red("Не найдено")
	}
}

func createAccount(vault *account.Vault) {
	login := scanTemplate("Введите логин: ")
	password := scanTemplate("Введите пароль: ")
	url := scanTemplate("Введите URL: ")

	accountOne, err := account.NewAccount(login, url, password)

	if err != nil {
		fmt.Println(err)
		return
	}
	vault.AddAccount(*accountOne)
}

func scanTemplate(template string) string {
	color.Cyan(template)
	var res string
	fmt.Scan(&res)
	return res
}
