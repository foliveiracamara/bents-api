package inmemory

import (
	"github.com/foliveiracamara/bents-api/application/entity"
	"github.com/foliveiracamara/bents-api/configuration/apperr"
)

type InMemoryEateryAdapter struct {
	eateries []*entity.Eatery
}

func NewInMemoryEateryAdapter() *InMemoryEateryAdapter {
	return &InMemoryEateryAdapter{
		eateries: []*entity.Eatery{},
	}
}

var eateries []*entity.Eatery

func (ea *InMemoryEateryAdapter) FindEatery(uuid string) (e *entity.Eatery, err *apperr.AppErr) {
	for _, item := range eateries {
		if item.UUID == uuid {
			return item, nil
		}
	}

	return nil, apperr.NewNotFoundError("Eatery non existent.")
}

func (ea *InMemoryEateryAdapter) CreateEatery(eatery *entity.Eatery) (err *apperr.AppErr) {
	eateries = append(eateries, eatery)

	return nil
}

func (ea *InMemoryEateryAdapter) FindEateries(searches []string) ([]*entity.Eatery, *apperr.AppErr) {
	// eateries = append(eateries, eatery)

	return nil, nil
}

// func (ea *InMemoryEateryAdapter) FindEateryByLocation(location string) (e *entity.Eatery, err *apperr.AppErr) {
// 	for _, item := range eateries {
// 		if item.Location == location {
// 			return item, nil
// 		}
// 	}

// 	return nil, apperr.NewNotFoundError("Eatery non existent.")
// }
