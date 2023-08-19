package auth

import (
	"errors"
	"fmt"
	"github.com/Shemistan/Lesson_3/utils"
	"os"
)

var users = make([]string, 0, 10)
var CurrentUser = ""

func Register(user string) (string, error) {
	CurrentUser = ""
	if utils.UserExists(users, user) {
		return "Этот пользователь уже существует", errors.New("user already exists")
	}
	users = append(users, user)

	message := fmt.Sprintf("Пользователь %s успешно добавлен", user)
	return message, nil
}

func Auth(user string) (string, error) {
	CurrentUser = ""
	if utils.UserExists(users, user) {
		CurrentUser = user
		return "Добро пожаловать в магазин", nil
	} else {
		return "Вы не зарегистрированы", errors.New("user not found")
	}
}

func Exit() {
	os.Exit(0)
}
