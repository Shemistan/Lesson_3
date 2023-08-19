package utils

import "fmt"

const (
	Reg        = "reg"
	Auth       = "auth"
	AddProduct = "add_product"
	Order      = "order"
	Cancel     = "cancel"
	Exit       = "exit"
)

const (
	regDescription        = Reg + " - регистрация нового профиля"
	authDescription       = Auth + " - авторизация в существующий профиль"
	addProductDescription = AddProduct + " - добавить продукт в корзину"
	orderDescription      = Order + " - создаёт чек и очищает корзину"
	cancelDescription     = Cancel + " - очищает корзину"
	exitDescription       = Exit + " - выход из приложения"
)

func showCommands() {
	fmt.Println("Введите команду") // Сделать красивый вывод, вывести список команд на этом шаге
	fmt.Println(regDescription)
	fmt.Println(authDescription)
	fmt.Println(addProductDescription)
	fmt.Println(orderDescription)
	fmt.Println(cancelDescription)
	fmt.Println(exitDescription)
}
