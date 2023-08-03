package main

import (
	"fmt"

	"github.com/Shemistan/Lesson_3/store"
)

const (
	exit       = "exit"
	auth       = "auth"
	reg        = "reg"
	addProduct = "add_product"
	order      = "order"
)

func main() {
	var command string
	var checkAuth bool
	userList := []string{"user1_password1", "user2_pass2"}
	productList := make([]string, 0, 10)

	listOperations()

outerLoop:
	for command != exit {
		fmt.Println("Введите команду")
		_, err := fmt.Scan(&command)
		if err != nil {
			return
		}

		switch command {
		case exit:
			break
		case reg:
			fmt.Println("Введите логин и пароль в таком виде login_passwor")
			_, err := fmt.Scan(&command)
			if err != nil {
				return
			}
			res := store.SaveUser(command, &userList)
			if !res {
				continue outerLoop
			}
			checkAuth = true
		case auth:
			fmt.Println("Введите логин и пароль в таком виде login_passwor")
			_, err := fmt.Scan(&command)
			if err != nil {
				return
			}
			checkAuth = store.CheckAuth(&userList, command)

		case addProduct:
			if !checkAuth {
				fmt.Println("Вы не авторизованны")
				continue outerLoop
			}
			fmt.Println("Добавьте продукт:")
			_, err := fmt.Scan(&command)
			if err != nil {
				return
			}
			res := store.AddProduct(command, &productList)
			if !res {
				continue outerLoop
			}
		case order:
			if !checkAuth {
				fmt.Println("Вы не авторизованны")
				continue outerLoop
			}
			res := store.Order(&productList)
			if !res {
				continue outerLoop
			}
		}
	}
}

func listOperations() {
	fmt.Println("Список команд:")
	fmt.Printf("Команда %s - выход из приложения\n", exit)
	fmt.Printf("Команда %s - авторизацияя\n", auth)
	fmt.Printf("Команда %s - регистрация\n", reg)
	fmt.Printf("Команда %s - добавление продукта (необходимо предварительно авторизоваться)\n", addProduct)
	fmt.Printf("Команда %s - покупка и очистка корзины (необходимо предварительно авторизоваться)\n", order)
}
