package auth

import (
	"fmt"
	"github.com/Shemistan/Lesson_3/utils"
	"os"
)

var users = make([]string, 0, 10)
var CurrentUser = ""

func Register(user string) string {
	CurrentUser = ""
	if utils.UserExists(users, user) {
		return "Этот пользователь уже существует"
	}
	users = append(users, user)

	message := fmt.Sprintf("Пользователь %s успешно добавлен", user)
	return message
}

func Auth(user string) string {
	CurrentUser = ""
	if utils.UserExists(users, user) {
		CurrentUser = user
		return "Добро пожаловать в магазин"
	} else {
		return "Вы не зарегистрированы"
	}
}

func Exit() {
	os.Exit(0)
}

func LogOut() {
	CurrentUser = ""
}

func IsAuthorized() bool {
	return CurrentUser != ""
}
