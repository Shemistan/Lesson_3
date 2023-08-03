package store

import "fmt"

func CheckAuth(userList *[]string, user string) bool {
	check := checkAuth(*userList, user)
	if check {
		fmt.Println("Добро пожаловать в магазин")
		return true
	}
	fmt.Println("Вы не зарегистрированны")
	return false
}

func SaveUser(user string, userList *[]string) bool {
	check := checkAuth(*userList, user)
	if check {
		fmt.Println("Пользователь уже существует")
		return false
	}

	*userList = append(*userList, user)
	message := fmt.Sprintf("Пользователь %s успешно добавлен", user)
	fmt.Println(message)
	fmt.Println(userList)
	return true
}

func checkAuth(userList []string, user string) bool {
	for _, v := range userList {
		if v == user {
			return true
		}
	}
	return false
}
