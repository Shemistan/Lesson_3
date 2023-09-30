package services

import "fmt"

var users = make(map[string]struct{})

func Register(u string) {
	if _, ok := users[u]; ok {
		fmt.Println("Пользователь уже существует!")
		return
	}

	users[u] = struct{}{}
	fmt.Println("Пользователь успешно добавлен")
}

func Login(u string) {
	if _, ok := users[u]; ok {
		fmt.Printf("Добро пожаловать, %s! \n", u)
		return
	}

	fmt.Println("Такой пользователь не существует")
}