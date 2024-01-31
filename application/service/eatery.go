package service

import (
	"github.com/foliveiracamara/bents-api/application/entity"
	port "github.com/foliveiracamara/bents-api/application/port/driven"
	"github.com/foliveiracamara/bents-api/configuration/apperr"
	security "github.com/foliveiracamara/bents-api/configuration/security/encryption"
	"github.com/rs/zerolog/log"
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

func (us *EateryService) GetEatery(name string) (eat []*entity.Eatery, err *apperr.AppErr) {
	if name == "" {
		return eat, apperr.NewBadRequestError("Eatery name cant be empty.")
	}

	res, err := us.EateryPort.FindEateryByName(name)
	if err != nil {
		return eat, err
	}

	return res, nil
}

func (us *EateryService) FindEateries(filters map[string]interface{}) (e []*entity.Eatery, err *apperr.AppErr) {
	var eateries []*entity.Eatery
	for filter, filterValue := range filters {
		if filter == "rank" {
			rankFilter, ok := filterValue.(int)
			if !ok {
				return nil, apperr.NewInternalServerError("Invalid filter type.")
			}

			eatery, err := us.EateryPort.FindEateriesByRank(rankFilter)
			if err != nil {
				return nil, err
			}

			for _, eat := range eatery {
				eateries = append(eateries, eat)
			}
		} else if filter == "category" {
			categoryFilter, ok := filterValue.(string)
			if !ok {
				return nil, apperr.NewInternalServerError("Invalid filter type.")
			}
			eatery, err := us.EateryPort.FindEateriesByCategory(categoryFilter)
			if err != nil {
				return nil, err
			}

			for _, eat := range eatery {
				eateries = append(eateries, eat)
			}
		}
	}

	if len(eateries) == 0 {
		return nil, apperr.NewNotFoundError("No eateries found.")
	}

	log.Info().
		Str("journey", "eateryService.FindEateries").
		Msg("Eateries found successfully.")

	return eateries, nil
}
