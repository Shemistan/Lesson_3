package main

import (
	"fmt"
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
	var userList []string
	productList := make([]string, 0, 10)
	authenticated := false

	for command != exit {
		showCommands(authenticated)
		fmt.Print("Введите номер команды: ")
		fmt.Scan(&command)

		switch command {
		case "1":
			fmt.Println("Программа завершена.")
			return
		case "2":
			registerUser(&userList)
		case "3":
			authenticated = authenticateUser(&userList)
		case "4":
			addProductToCart(authenticated, &productList)
		case "5":
			processOrder(authenticated, &productList)
		default:
			fmt.Println("Неверная команда. Пожалуйста, выберите команду из списка.")
		}
	}
}

func showCommands(authenticated bool) {
	fmt.Println("Выберите команду:")
	fmt.Println("1. " + exit + ": Выход из программы")
	fmt.Println("2. " + reg + ": Регистрация нового пользователя")
	fmt.Println("3. " + auth + ": Авторизация пользователя")
	if authenticated {
		fmt.Println("4. " + addProduct + ": добавить продукт в корзину")
		fmt.Println("5. " + order + ": показать корзину и оформить заказ")
	}
}

func registerUser(userList *[]string) {
	var command string
	fmt.Print("Введите логин и пароль в формате login_password: ")
	fmt.Scan(&command)
	if userExists(*userList, command) {
		fmt.Println("Пользователь уже существует.")
	} else {
		*userList = append(*userList, command)
		fmt.Printf("Пользователь %s успешно добавлен.\n", command)
	}
}

func authenticateUser(userList *[]string) bool {
	var command string
	fmt.Print("Введите логин и пароль в формате login_password: ")
	fmt.Scan(&command)
	if userExists(*userList, command) {
		fmt.Println("Добро пожаловать в магазин!")
		return true
	}
	fmt.Println("Вы не зарегистрированы.")
	return false
}

func addProductToCart(authenticated bool, productList *[]string) {
	if authenticated {
		var command string
		fmt.Print("Введите название продукта: ")
		fmt.Scan(&command)
		*productList = append(*productList, command)
		fmt.Printf("Продукт %s успешно добавлен в корзину.\n", command)
	} else {
		fmt.Println("Вы не авторизованы.")
	}
}

func processOrder(authenticated bool, productList *[]string) {
	if authenticated {
		if len(*productList) == 0 {
			fmt.Println("Ваша корзина пуста.")
		} else {
			fmt.Println("Ваша корзина:")
			for _, v := range *productList {
				fmt.Println(v)
			}
			fmt.Println("Заказ успешно оформлен.")
			*productList = make([]string, 0, 10)
		}
	} else {
		fmt.Println("Вы не авторизованы.")
	}
}

func userExists(userList []string, user string) bool {
	for _, v := range userList {
		if v == user {
			return true
		}
	}
	return false
}
