package repositories

import (
	"github.com/Shemistan/Lesson_3/errors"
	"github.com/Shemistan/Lesson_3/models"
)

type UserRepository struct {
	dbHandler *models.Database
}

func NewUserRepository(dbHandler *models.Database) *UserRepository {
	return &UserRepository{dbHandler: dbHandler}
}

func (ur UserRepository) GetUser(login, password string) (models.User, error) {
	if len(ur.dbHandler.Users) == 0 {
		return models.User{}, errors.CollectionEmpty
	}

	for _, user := range ur.dbHandler.Users {
		if user.Login == login && user.Password == password {
			return user, nil
		}
	}

	return models.User{}, errors.NotExists
}

func (ur UserRepository) AddUser(user models.User) error {
	user.Id = len(ur.dbHandler.Users)
	ur.dbHandler.Users = append(ur.dbHandler.Users, user)
	return nil
}

func (ur UserRepository) UpdateUser(id int, user models.User) error {
	if len(ur.dbHandler.Users) == 0 {
		return errors.CollectionEmpty
	}

	ur.dbHandler.Users[id] = user
	return nil
}
