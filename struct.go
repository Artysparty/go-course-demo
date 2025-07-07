package main

import (
	"demo/app-1/account"
	"demo/app-1/files"
	"demo/app-1/output"
	"fmt"

	"github.com/fatih/color"
)

func structures() {
	fmt.Println("Менеджер паролей")
	vault := account.NewVault(files.NewJsonDB("data.json"))
	// vault := account.NewVault(cloud.NewCloudDB("https://g.ru"))

Menu:
	for {
		variant := scanTemplate([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите вариант",
		})

		switch variant {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		case "4":
			break Menu
		}
	}
}

func findAccount(vault *account.VaultWithDB) {
	url := scanTemplate([]string{"Введите url для поиска"})

	accounts := vault.FindAccountByUrl(url)

	if len(accounts) == 0 {
		output.PrintError("Аккаунтов не найдено")
	}

	for _, account := range accounts {
		account.PrintAccount()
	}
}

func deleteAccount(vault *account.VaultWithDB) {
	url := scanTemplate([]string{"Введите url для удаления"})

	isDeleted := vault.DeleteAccountByUrl(url)

	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Аккаунтов не найдено")
	}
}

func createAccount(vault *account.VaultWithDB) {
	login := scanTemplate([]string{"Введите логин: "})
	password := scanTemplate([]string{"Введите пароль: "})
	url := scanTemplate([]string{"Введите URL: "})

	accountOne, err := account.NewAccount(login, url, password)

	if err != nil {
		fmt.Println(err)
		return
	}
	vault.AddAccount(*accountOne)
}

func scanTemplate[T any](templates []T) string {
	for i, tpl := range templates {
		if i == len(templates)-1 {
			fmt.Printf("%v: ", tpl)
		} else {
			fmt.Println(tpl)
		}
	}
	var res string
	fmt.Scan(&res)
	return res
}
