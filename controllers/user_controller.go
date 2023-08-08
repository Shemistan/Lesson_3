package controllers

import (
	"github.com/Shemistan/Lesson_3/errors"
	"github.com/Shemistan/Lesson_3/models"
	"github.com/Shemistan/Lesson_3/services"
)

type UserController struct {
	usersService *services.UserService
}

func NewUserController(usersService *services.UserService) *UserController {
	return &UserController{usersService: usersService}
}

func (uc UserController) GetUser(login, password string) (models.User, error) {
	if login == "" || password == "" {
		return models.User{}, errors.MissingParams
	}

	return uc.usersService.GetUser(login, password)
}

func (uc UserController) AddUser(user models.User) error {
	if user.Login == "" || user.Password == "" || user.FullName == "" || user.PhoneNumber == "" {
		return errors.MissingParams
	}

	return uc.usersService.AddUser(user)
}

func (uc UserController) UpdateUser(id int, user models.User) error {
	if user.Login == "" || user.Password == "" || user.FullName == "" || user.PhoneNumber == "" {
		return errors.MissingParams
	}

	return uc.usersService.UpdateUser(id, user)
}

func (uc UserController) FillUp(user models.User, amount float64) error {

	if amount <= 0 {
		return errors.InvalidArg
	}

	return uc.usersService.BalanceOperations(user, amount)
}

func (uc UserController) Withdraw(user models.User, amount float64) error {
	if amount <= 0 {
		return errors.InvalidArg
	}

	if amount > user.Balance {
		return errors.InvalidArg
	}

	return uc.usersService.BalanceOperations(user, -amount)
}
