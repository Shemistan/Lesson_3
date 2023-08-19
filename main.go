package main

import (
	"fmt"
	"github.com/Shemistan/Lesson_3/auth"
	"github.com/Shemistan/Lesson_3/constants"
	"github.com/Shemistan/Lesson_3/store"
	"github.com/Shemistan/Lesson_3/utils"
)

var command string

func main() {
	for command != utils.Exit {
		handleCommand(command)
	}
}

func handleCommand(command string) {
	if command == "" {
		utils.ShowCommands(auth.CurrentUser)
		utils.Read(&command)
	}

	switch command {
	case utils.Reg:
		handleReg()
	case utils.Auth:
		handAuth()
	case utils.AddProduct:
		handleAddProduct()
	case utils.Order:
		handleOrder()
	case utils.Cancel:
		handleCancel()
	case utils.Exit:
		handleExit()
	default:
		command = ""
	}
}

func handleReg() {
	fmt.Println(constants.ENTER_LOGIN_AND_PASSWORD)
	var user string
	utils.Read(&user)

	message := auth.Register(user)
	fmt.Println(message)
	command = ""
}

func handAuth() {
	fmt.Println(constants.ENTER_LOGIN_AND_PASSWORD)
	var user string
	utils.Read(&user)

	message := auth.Auth(user)
	fmt.Println(message)
	command = ""
}

func handleAddProduct() {
	if auth.IsAuthorized() {
		fmt.Println(constants.ENTER_PRODUCT)
		var product string
		utils.Read(&product)

		message := store.AddProduct(auth.CurrentUser, product)
		fmt.Println(message)
		command = ""
	} else {
		fmt.Println(constants.NOT_AUTHORIZED)
		command = utils.Auth
	}
}

func handleOrder() {
	if auth.IsAuthorized() {
		message, err := store.Order(auth.CurrentUser)
		fmt.Println(message)
		if err != nil {
			command = utils.AddProduct
		} else {
			command = ""
		}
	} else {
		fmt.Println(constants.NOT_AUTHORIZED)
		command = utils.Auth
	}
}

func handleCancel() {
	if auth.IsAuthorized() {
		message := store.Cancel(auth.CurrentUser)
		fmt.Println(message)
		auth.LogOut()
		command = ""
	} else {
		fmt.Println(constants.NOT_AUTHORIZED)
		command = utils.Auth
	}
}

func handleExit() {
	auth.LogOut()
	auth.Exit()
}
