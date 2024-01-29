package service

import (
	"fmt"
	"regexp"

	"github.com/foliveiracamara/bents-api/application/entity"
	port "github.com/foliveiracamara/bents-api/application/port/driven"
	"github.com/foliveiracamara/bents-api/configuration/apperr"
	security "github.com/foliveiracamara/bents-api/configuration/security/encryption"
)

type EateryService struct {
	EateryPort port.EateryPort
	Encrypt    *security.Encryption
}

func NewEateryService(eatPort port.EateryPort) *EateryService {
	return &EateryService{
		EateryPort: eatPort,
	}
}

func (es *EateryService) CreateEatery(user *entity.Eatery) (e *entity.Eatery, err *apperr.AppErr) {
	pwd, err := es.Encrypt.EncryptPassword(user.Password)
	if err != nil {
		return e, err
	}
	user.Password = pwd
	
	err = es.EateryPort.CreateEatery(user)
	if err != nil {
		appErr := apperr.NewInternalServerError("Error creating user.")
		return e, appErr
	}

	return user, nil
}

func (us *EateryService) GetEatery(uuid string) (u *entity.Eatery, err *apperr.AppErr) {
	if uuid == "" || len(uuid) > 36 || len(uuid) < 36 {
		return u, apperr.NewBadRequestError("Invalid UUID.")
	}

	uuidRegex := `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`
	match, _ := regexp.MatchString(uuidRegex, uuid)
	if !match {
		return u, apperr.NewBadRequestError("Invalid UUID.")
	}

	u, err = us.EateryPort.FindEatery(uuid)
	if err != nil {
		return u, err
	}

	return u, nil
}

func (us *EateryService) FindEateries(filters ...string) (e []*entity.Eatery, err *apperr.AppErr) {
	fmt.Println("filters in service: ", filters)

	// var searches []string

	// list, err := us.EateryPort.FindEateries()
	
	return e, nil
}
