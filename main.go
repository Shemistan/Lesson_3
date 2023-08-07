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
	fmt.Println(
		"Доступные комманды: \n",
		"reg         зарегестрироваться в приложении\n",
		"login       войти в аккаунт\n",
		"exit        выйти из приложения\n",
		"add_product добавить продукт в корзину\n",
		"order       совершить покупку вывести на консоль все купленные продукты",
	)
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
