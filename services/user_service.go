package services

import (
	"github.com/Shemistan/Lesson_3/models"
	"github.com/Shemistan/Lesson_3/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (us UserService) GetUser(login, password string) (models.User, error) {
	return us.userRepository.GetUser(login, password)
}

func (us UserService) AddUser(user models.User) error {
	return us.userRepository.AddUser(user)
}

func (us UserService) UpdateUser(id int, user models.User) error {
	return us.userRepository.UpdateUser(id, user)
}

func (us UserService) BalanceOperations(user models.User, amount float64) error {
	user.Balance += amount
	return us.userRepository.UpdateUser(user.Id, user)
}
