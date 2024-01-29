package port

import (
	"github.com/foliveiracamara/bents-api/application/entity"
	"github.com/foliveiracamara/bents-api/configuration/apperr"
)

type UserService interface {
	GetUser(uuid string) (*entity.User, *apperr.AppErr)
	CreateUser(user *entity.User) (u *entity.User, err *apperr.AppErr)
	LoginUser(email, password string) (*apperr.AppErr)
}

type EateryService interface {
	GetEatery(uuid string) (*entity.Eatery, *apperr.AppErr)
	CreateEatery(user *entity.Eatery) (u *entity.Eatery, err *apperr.AppErr)
	FindEateries(filters ...string) ([]*entity.Eatery, *apperr.AppErr)
}