package main

import "fmt"

type account struct {
	login    string
	password string
	url      string
}

func structures() {
	login := scanTemplate("Введите логин: ")
	password := scanTemplate("Введите пароль: ")
	url := scanTemplate("Введите URL: ")

	accountOne := account{
		login,
		password,
		url,
	}

	printAccount(&accountOne)
}

func scanTemplate(template string) string {
	fmt.Print(template)
	var res string
	fmt.Scan(&res)
	return res
}

func printAccount(acc *account) {
	fmt.Println(acc.login, acc.password, acc.url)
}
