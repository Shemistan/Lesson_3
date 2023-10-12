package auth

import (
	"fmt"
)

func Register(user_pass string, userList []string) ([]string, error) {
	for _, v := range userList {
		if v == user_pass {
			return userList, fmt.Errorf("Пользователь %s уже зарегистрирован", user_pass)
		}
	}
	userList = append(userList, user_pass)

	message := fmt.Sprintf("Пользователь %s успешно добавлен", user_pass)
	fmt.Println(message)

	return userList, nil
}

func Login(user_pass string, userList []string) error {
	for _, v := range userList {
		if v == user_pass {
			fmt.Println("Добро пожаловать в магазин")
			return nil
		}
	}
	return fmt.Errorf("Пользователь %s не найден", user_pass)
}
