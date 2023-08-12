package store

import "fmt"

func StartPrint() {
	fmt.Println("Введите команду, одну из exit/auth/reg/add_product/order")
}

// добавляем нового пользователя
func Registration(userList *[]string) bool {
	var command string
	fmt.Println("Введите логин и пароль в таком виде login_password")
	fmt.Scan(&command)
	*userList = append(*userList, command)

	message := fmt.Sprintf("Пользователь %s успешно добавлен", command)
	fmt.Println(message)
	return true
}

// Аутентификация пользователя, если не удалась - возвращаем неуспешный флаг
func Auth(flagReg bool, userList *[]string) bool {
	var flagAuth bool
	var command string
	if flagReg == true {
		fmt.Println("Добро пожаловать в магазин")
		flagAuth = true
	} else {
		fmt.Println("Введите логин и пароль в таком виде login_password")
		fmt.Scan(&command)

		for _, v := range *userList {
			if v == command {
				fmt.Println("Добро пожаловать в магазин")
				flagAuth = true
			}
		}
		if flagAuth != true {
			fmt.Println("Вы не зарегистрированны или неверный пароль")
			flagAuth = false
		}
	}
	return flagAuth
}

// если авторизованы - добавляем товары
func ProductAdd(flag_auth bool, productList *[]string) {
	var command string
	if flag_auth == true {
		fmt.Println("Какой товар хотите добавить?")
		fmt.Scan(&command)
		*productList = append(*productList, command)
		fmt.Println("Следующие товары в корзине", *productList)
	} else {
		fmt.Println("Вы не авторизованы")
	}
}

// заказываем если есть товары и прошли авторизацию
func Order(flag_auth bool, productList *[]string) {
	if flag_auth == true && len(*productList) > 0 {
		fmt.Println("Товары из корзины заказаны,", *productList, ", корзина очищена")
		*productList = nil
	} else if flag_auth == false {
		fmt.Println("Вы не авторизованы")
	} else {
		fmt.Println("Нет товаров в корзине, добавьте товары через add_product")
	}
}
