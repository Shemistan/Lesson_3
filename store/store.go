package store

import (
	"errors"
	"fmt"
)

var Products = make(map[string][]string)

func AddProduct(user, product string) string {
	Products[user] = append(Products[user], product)
	return fmt.Sprintf("%s добавлен в корзину", product)
}

func Order(user string) (string, error) {
	usersProducts := Products[user]
	if usersProducts == nil || len(usersProducts) == 0 {
		return "Корзина пуста", errors.New("cart is empty")
	} else {
		message := fmt.Sprintf("Чек (%s):", user)
		for i := range usersProducts {
			message += fmt.Sprint(i+1, usersProducts[i])
		}
		Products[user] = []string{}
		return message, nil
	}
}

func Cancel(user string) (string, error) {
	usersProducts := Products[user]
	if usersProducts == nil || len(usersProducts) == 0 {
		return "Корзина пуста", errors.New("cart is empty")
	} else {
		Products[user] = []string{}
		return "Корзина очищена", nil
	}
}
