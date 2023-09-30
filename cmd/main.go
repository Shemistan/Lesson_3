package main

import (
	"fmt"

	"github.com/Shemistan/Lesson_3/services"
)

const (
	exit       = "exit"
	auth       = "auth"
	reg        = "reg"
	addProduct = "add_product"
	order      = "order"
)

func main() {
	showHints()

outer:
	for {
		var command string
		fmt.Scan(&command)

		switch command {
		case exit:
			break outer
		case reg:
			fmt.Println("\n Введите логин и пароль в таком виде login_password")
			fmt.Scan(&command)
			services.Register(command)
		case auth:
			fmt.Println("\n Введите логин и пароль в таком виде login_password")
			fmt.Scan(&command)
			services.Login(command)
		case addProduct:
			fmt.Println("\n Введите название товара")
			fmt.Scan(&command)
			services.AddProduct(command)
		case order:
			fmt.Println("\n Список продуктов, которые вы приобрели:")
			services.Order()
		default:
			fmt.Println("\n Такой команды не существует!")
		}

	}
}

func showHints() {
	fmt.Println("Введите команду:")
	fmt.Printf("%s - %s\n", exit, "Выход")
	fmt.Printf("%s - %s\n", reg, "Регистрация")
	fmt.Printf("%s - %s\n", auth, "Авторизация")
	fmt.Printf("%s - %s\n", addProduct, "Добавить продукт в корзину")
	fmt.Printf("%s - %s\n", order, "Завершить покупку")
}
