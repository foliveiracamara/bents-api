package port

import (
	"github.com/foliveiracamara/bents-api/application/entity"
	"github.com/foliveiracamara/bents-api/configuration/apperr"
)

// @TODO: change Get to Find
type UserPort interface {
	GetUser(uuid string) (*entity.User, *apperr.AppErr)
	CreateUser(user *entity.User) (err *apperr.AppErr)
	FindUserByEmailAndPassword(email, password string) (*entity.User, *apperr.AppErr)
	FindUserPasswordByEmail(email string) (pwd string, err *apperr.AppErr)
}

type EateryPort interface {
	FindEatery(uuid string) (*entity.Eatery, *apperr.AppErr)
	CreateEatery(user *entity.Eatery) (err *apperr.AppErr)
	FindEateries(searches []string) ([]*entity.Eatery, *apperr.AppErr)
}
