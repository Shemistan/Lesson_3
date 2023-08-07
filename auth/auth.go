package auth

import (
	"fmt"
)

var users []string

func Register(user string) {
	if contains(users, user) {
		fmt.Println("Пользователь уже зарегестрирован")
		return
	}
	users = append(users, user)
	fmt.Printf("Пользователь %s успешно зарегестрирован \n", user)
}

func Login(user string) {
	if contains(users, user) {
		fmt.Printf("%s Добро пожаловать \n", user)
		return
	}
	fmt.Println("Пользователя нету в системе")
}

func contains(slice []string, element string) bool {
	for _, sliceElement := range slice {
		if sliceElement == element {
			return true
		}
	}
	return false
}
