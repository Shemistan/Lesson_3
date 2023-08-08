package main

import "fmt"

const (
	exit = "exit"
	auth = "auth"
	reg  = "reg"
	help = "help"
	add_product = "add_product"
	order = "order"

)

func checkExist(List []string, item string) bool {
	for _, v := range List {
		if v == item {
			return true
		}
	}

	return false
}

func help_func() {
	fmt.Println("\n",
	reg, "- Регистрация пользователя.\n",
	auth, "- Авторизация.\n",
	exit, "- Закончить работу.\n",
	add_product, "- Добавить продукт в корзину.\n",
	order, "- Заказать товары из корзины.\n")
}

func reg_func(userList []string) []string {
	var login_password string

	fmt.Println("Введите логин и пароль в формате login_password:")
	fmt.Scan(&login_password)

	if checkExist(userList, login_password) == true {
		fmt.Println("Такой пользователь уже существует")
	} else {
		userList = append(userList, login_password)
		message := fmt.Sprintf("Пользователь %s успешно добавлен", login_password)
		fmt.Println(message)
	}

	return userList
}

func auth_func(userList []string) {
	var login_password string

	fmt.Println("Введите логин и пароль в таком виде login_password")
	fmt.Scan(&login_password)

	if checkExist(userList, login_password) == true {
		fmt.Println("Добро пожаловать в магазин")
	} else {
		fmt.Println("Вы не зарегистрированы")
	}
}

func add_product_func(productList []string) []string  {
	var product string

	fmt.Println("Введите наименование товара:")
	fmt.Scan(&product)

	productList = append(productList, product)
	message := fmt.Sprintf("Товар %s успешно добавлен", product)
	fmt.Println(message)

	return productList
}

func order_func(productList []string){
	if len(productList) == 0 {
		fmt.Println("Корзина пуста")
	} else {
		fmt.Println("Спасибо за заказ. Ваш список товаров:")
		for _, v := range productList {
			fmt.Println(v)
		}
	}
}

func default_func() {
	fmt.Println("Введена неправильная команда.")
}

func main() {
	var command string
	userList := []string{"user1_password1", "user2_password2"}
	productList := make([]string, 0, 10)

	_ = productList
	for command != exit {
		fmt.Println("Введите команду. Введите help для получения списка команд.")
		fmt.Scan(&command)

		switch command {
		case help:
			help_func()
		case reg:
			userList = reg_func(userList)
		case auth:
			auth_func(userList)
		case add_product:
			productList = add_product_func(productList)
		case order:
			order_func(productList)
			productList = productList[:0]
		case exit:
			break
		default:
			default_func()
		}
	}
}

// Реализовать следующие АПИ
// add_product - добавить продукт который вводиться с консоли в productList
// order - выводит сообщение что вы купили и очищает корзину
