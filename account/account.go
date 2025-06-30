package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"unicode/utf8"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890?!-.,_")

// Структура account
type Account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Url      string `json:"url"`
}

// Методы структуры account
func (acc *Account) PrintAccount() {
	fmt.Println(acc.Login, acc.Password, acc.Url)
}

func (acc *Account) ToBytes() ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(res)
}

// Функция - конструктор
func NewAccount(login, urlString, password string) (*Account, error) {
	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("Неверный формат URL")
	}

	account := Account{
		Login:    login,
		Url:      urlString,
		Password: password,
	}

	if utf8.RuneCountInString(account.Password) == 0 {
		account.generatePassword(12)
	}

	return &account, nil
}
