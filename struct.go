package main

import (
	"demo/app-1/account"
	"demo/app-1/encrypter"
	"demo/app-1/files"
	"demo/app-1/output"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithDB){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

var menuVariants = []string{
	"1. Создать аккаунт",
	"2. Найти аккаунт по Url",
	"3. Найти аккаунт по логину",
	"4. Удалить аккаунт",
	"5. Выход",
	"Выберите вариант",
}

func menuCounter() func() {
	i := 0
	return func() {
		i++
		fmt.Println(i)
	}
}

func structures() {
	fmt.Println("Менеджер паролей")
	vault := account.NewVault(files.NewJsonDB("data.vault"), *encrypter.NewEncrypter())
	counter := menuCounter()

	error := godotenv.Load()

	if error != nil {
		output.PrintError("Не удалось найти .env")
	}

	res := os.Getenv("VAR")
	fmt.Println(res)

Menu:
	for {
		counter()
		variant := scanTemplate(
			menuVariants...,
		)

		menuFunc := menu[variant]

		if menuFunc == nil {
			break Menu
		}

		menuFunc(vault)
	}
}

func findAccountByUrl(vault *account.VaultWithDB) {
	url := scanTemplate("Введите url для поиска")

	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})

	if len(accounts) == 0 {
		output.PrintError("Аккаунтов не найдено")
	}

	for _, account := range accounts {
		account.PrintAccount()
	}
}

func findAccountByLogin(vault *account.VaultWithDB) {
	login := scanTemplate("Введите login для поиска")

	accounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})

	if len(accounts) == 0 {
		output.PrintError("Аккаунтов не найдено")
	}

	for _, account := range accounts {
		account.PrintAccount()
	}
}

func deleteAccount(vault *account.VaultWithDB) {
	url := scanTemplate("Введите url для удаления")

	isDeleted := vault.DeleteAccountByUrl(url)

	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Аккаунтов не найдено")
	}
}

func createAccount(vault *account.VaultWithDB) {
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

func scanTemplate(templates ...string) string {
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
