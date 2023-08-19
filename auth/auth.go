package auth

import (
	"fmt"
	"github.com/Shemistan/Lesson_3/constants"
	"github.com/Shemistan/Lesson_3/utils"
	"os"
)

var users = make([]string, 0, 10)
var CurrentUser = ""

func Register(user string) string {
	CurrentUser = ""
	if utils.UserExists(users, user) {
		return constants.USER_ALREADY_EXISTS
	}
	users = append(users, user)

	message := fmt.Sprintf(constants.USER_ADDED, user)
	return message
}

func Auth(user string) string {
	CurrentUser = ""
	if utils.UserExists(users, user) {
		CurrentUser = user
		return constants.WELCOME
	} else {
		return constants.NOT_AUTHORIZED
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
