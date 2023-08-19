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

func ShowCommands(user string) {
	fmt.Println("Введите команду")
	if user == "" {
		fmt.Println(regDescription)
		fmt.Println(authDescription)
		fmt.Println(exitDescription)
	} else {
		fmt.Println(addProductDescription)
		fmt.Println(orderDescription)
		fmt.Println(cancelDescription)
	}
}
