package main

import (
	"fmt"
	"github.com/Shemistan/Lesson_3/auth"
	"github.com/Shemistan/Lesson_3/store"
)

const (
	reg        = "reg"
	login      = "login"
	exit       = "exit"
	addProduct = "add_product"
	order      = "order"
)

func main() {
	fmt.Println("Доступные комманды: ")

	showCommandHint(reg, "зарегестрироваться в приложении")
	showCommandHint(login, "войти в аккаунт")
	showCommandHint(exit, "выйти из приложения")
	showCommandHint(addProduct, "добавить продукт в корзину")
	showCommandHint(order, "совершить покупку вывести на консоль все купленные продукты")

loop:
	for {
		var command string
		fmt.Scan(&command)

		switch command {
		case exit:
			fmt.Println("Выходим из приложения...")
			break loop
		case reg:
			fmt.Println("Введите данные пользователя в формате login_password")
			fmt.Scan(&command)
			auth.Register(command)
		case login:
			fmt.Println("Введите данные пользователя в формате login_password")
			fmt.Scan(&command)
			auth.Login(command)
		case addProduct:
			fmt.Println("Введите название товара который вы хотите добавить")
			fmt.Scan(&command)
			store.AddProduct(command)
		case order:
			fmt.Println(store.Order())
		}
	}
}

func showCommandHint(command string, hint string) {
	fmt.Printf("%s - %s\n", command, hint)
}

func increment() {
	defer fmt.Println("вызвался defer 1")
	defer fmt.Println("вызвался defer 2")
	defer fmt.Println("вызвался defer 3")
	fmt.Println("1")
	fmt.Println("2")
	var a map[int]int

	a[1] = 1

	fmt.Println("3")
}
