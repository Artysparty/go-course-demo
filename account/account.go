package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"
	"unicode/utf8"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890?!-.,_")

type Account struct {
	Login       string    `json:"login"`
	Password    string    `json:"password"`
	Url         string    `json:"url"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
}

func (acc *Account) PrintAccount() {
	color.Blue(acc.Login)
	color.Blue(acc.Password)
	color.Blue(acc.Url)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(res)
}

func NewAccount(login, urlString, password string) (*Account, error) {
	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("неверный формат URL")
	}

	account := Account{
		Login:       login,
		Url:         urlString,
		Password:    password,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}

	if utf8.RuneCountInString(account.Password) == 0 {
		account.generatePassword(12)
	}

	return &account, nil
}
