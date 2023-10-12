package main

import (
	"fmt"
	au "github.com/Shemistan/Lesson_3/auth"
	"github.com/Shemistan/Lesson_3/store"
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
	userList := []string{"user1_password1", "user2_password2"}
	productList := make([]string, 0, 10)

	for command != exit {
		initialCommands(&command)

		switch command {
		case exit:
			break
		case reg:
			input(&command, "Введите логин и пароль в таком виде login_password")
			userList, err := au.Register(command, userList)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(userList)
		case auth:
			input(&command, "Введите логин и пароль в таком виде login_password")
			err := au.Login(command, userList)
			if err != nil {
				fmt.Println(err)
			}
		case add_product:
			input(&command, "Введите название продукта")
			productList = store.AddProduct(productList, command)
			fmt.Println(productList)
		case order:
			productList = store.Order(productList)
		}
	}
}

func input(command *string, message string) {
	fmt.Println(message)
	fmt.Scan(command)
}

func initialCommands(command *string) {
	fmt.Println("Введите команду")
	fmt.Printf("%v - для авторизации \n", auth)
	fmt.Printf("%v - для регистрации \n", reg)
	fmt.Printf("%v - для добавления продукта \n", add_product)
	fmt.Printf("%v - для покупки \n", order)
	fmt.Printf("%v - для выхода \n", exit)
	fmt.Scan(command)
}
