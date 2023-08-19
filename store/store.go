package store

import (
	"errors"
	"fmt"
	"github.com/Shemistan/Lesson_3/constants"
)

var Products = make(map[string][]string)

func AddProduct(user, product string) string {
	Products[user] = append(Products[user], product)
	return fmt.Sprintf(constants.PRODUCT_ADDED, product)
}

func Order(user string) (string, error) {
	usersProducts := Products[user]
	if usersProducts == nil || len(usersProducts) == 0 {
		return constants.CART_IS_EMPTY, errors.New("cart is empty")
	} else {
		message := fmt.Sprintf(constants.CART_TITLE, user)
		for i := range usersProducts {
			message += fmt.Sprintf(constants.CART_ROW, i+1, usersProducts[i])
		}
		Products[user] = []string{}
		return message, nil
	}
}

func Cancel(user string) string {
	usersProducts := Products[user]
	if usersProducts == nil || len(usersProducts) == 0 {
		return constants.CART_IS_EMPTY
	} else {
		Products[user] = []string{}
		return constants.CART_EMPTIED
	}
}
