package inmemory

import (
	"github.com/foliveiracamara/bents-api/application/entity"
	"github.com/foliveiracamara/bents-api/configuration/apperr"
)

type InMemoryUserAdapter struct {
	users []*entity.User
}

func NewInMemoryUserAdapter() *InMemoryUserAdapter {
	return &InMemoryUserAdapter{
		users: []*entity.User{},
	}
}

var users []*entity.User

func (ua *InMemoryUserAdapter) GetUser(id string) (u *entity.User, err *apperr.AppErr) {
	for _, item := range users {
		if item.UUID == id {
			return item, nil
		}
	}

	return nil, apperr.NewNotFoundError("User non existent.")
}

func (ua *InMemoryUserAdapter) CreateUser(user *entity.User) (err *apperr.AppErr) {
	users = append(users, user)
	
	return nil
}

func (ua *InMemoryUserAdapter) FindUserByEmailAndPassword(email, password string) (u *entity.User, err *apperr.AppErr) {
	user := &entity.User{
		Email:    "fcamara@gmail.com",
		Password: "Passq0r4",
	}

	if email == user.Email && password == user.Password {
		return user, nil
	}

	return
}

func (ua *InMemoryUserAdapter) FindUserPasswordByEmail(email string) (pwd string, err *apperr.AppErr) {
	for _, item := range users {
		if item.Email == email {
			return item.Password, nil
		}
	}

	return "", apperr.NewNotFoundError("User non existent.")
}
