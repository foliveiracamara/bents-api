package service

import (
	"fmt"
	"regexp"

	"github.com/foliveiracamara/bents-api/application/entity"
	port "github.com/foliveiracamara/bents-api/application/port/driven"
	"github.com/foliveiracamara/bents-api/configuration/apperr"
	security "github.com/foliveiracamara/bents-api/configuration/security/encryption"
	"github.com/rs/zerolog/log"
)

type UserService struct {
	UserPort port.UserPort
	Encrypt  *security.Encryption
}

func NewUserService(userPort port.UserPort) *UserService {
	return &UserService{
		UserPort: userPort,
	}
}

func (us *UserService) GetUser(uuid string) (u *entity.User, err *apperr.AppErr) {
	if uuid == "" || len(uuid) > 36 || len(uuid) < 36 {
		return u, apperr.NewBadRequestError("Invalid UUID.")
	}

	uuidRegex := `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`
	match, _ := regexp.MatchString(uuidRegex, uuid)
	if !match {
		return u, apperr.NewBadRequestError("Invalid UUID.")
	}

	u, err = us.UserPort.GetUser(uuid)
	if err != nil {
		return u, err
	}

	return u, nil
}

func (us *UserService) CreateUser(user *entity.User) (u *entity.User, err *apperr.AppErr) {
	pwd, err := us.Encrypt.EncryptPassword(user.Password)
	if err != nil {
		return u, err
	}
	user.Password = pwd

	err = us.UserPort.CreateUser(user)
	if err != nil {
		appErr := apperr.NewInternalServerError("Error creating user.")
		return u, appErr
	}

	return user, nil
}

func (us *UserService) LoginUser(email string, passwordRequest string) (err *apperr.AppErr) {
	passwordDB, err := us.UserPort.FindUserPasswordByEmail(email)
	if err != nil {
		fmt.Println("Error finding user's password on database.")
		return err
	}

	err = us.Encrypt.CompareHashAndPassword(passwordDB, passwordRequest)
	if err != nil {
		log.Error().
			Str("journey", "userService.LoginUser").
			Msg(err.Error())
		appErr := apperr.NewUnauthorizedError("Invalid password.")
		return appErr
	}

	return nil
}
