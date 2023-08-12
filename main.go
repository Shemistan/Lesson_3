package main

import (
	"fmt"
	"github.com/Zzema/Lesson_3/store"
	//"github.com/Zzema/lesson_3/store"
)

const (
	exit        = "exit"
	auth        = "auth"
	reg         = "reg"
	add_product = "add_product"
	order       = "order"
)

func main() {
	var command string
	var flagAuth, flagReg bool //признак идентификации пользователя
	userList := []string{"user1_password1", "user2_password2"}
	productList := make([]string, 0, 10)

	//_ = productList
	for command != exit {
		store.StartPrint()
		fmt.Scan(&command)

		switch command {
		case exit:
			fmt.Println("exiting")
			break
		case reg:
			flagReg = store.Registration(&userList)

		case auth:
			flagAuth = store.Auth(flagReg, &userList)
		case add_product:
			store.ProductAdd(flagAuth, &productList)

		case order:
			store.Order(flagAuth, &productList)
		}
	}

}
