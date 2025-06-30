package main

import (
	"demo/app-1/account"
	"demo/app-1/files"
	"fmt"

	"github.com/fatih/color"
)

// Основной код
func structures() {
	login := scanTemplate("Введите логин: ")
	password := scanTemplate("Введите пароль: ")
	url := scanTemplate("Введите URL: ")

	accountOne, err := account.NewAccount(login, url, password)

	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := accountOne.ToBytes()

	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
		return
	}

	files.WriteFile(file, "data.json")

	// accountOne.PrintAccount()
}

func scanTemplate(template string) string {
	color.Cyan(template)
	var res string
	fmt.Scan(&res)
	return res
}
