package inmemory

import (
	"fmt"
	"strings"

	"github.com/foliveiracamara/bents-api/application/entity"
	"github.com/foliveiracamara/bents-api/configuration/apperr"
	"github.com/rs/zerolog/log"
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

func (ea *InMemoryEateryAdapter) FindEateryByName(name string) (e []*entity.Eatery, err *apperr.AppErr) {
	var res []*entity.Eatery

	for _, item := range eateries {
		if strings.Contains(item.Name, name) {
			fmt.Println(true)
			res = append(res, item)
		}
	}

	errMsg := fmt.Sprintf("Eateries not found for '%v' search.", name)
	if len(res) == 0 {
		return nil, apperr.NewNotFoundError(errMsg)
	}

	fmt.Println("eateries by name: ", res)

	return res, nil
}

func (ea *InMemoryEateryAdapter) CreateEatery(eatery *entity.Eatery) (err *apperr.AppErr) {
	eateries = append(eateries, eatery)

	log.Info().
		Str("journey", "eateryAdapter.CreateEatery").
		Msgf("Eatery '%s' created. rank %v", eatery.Name, eatery.Rank)

	return nil
}

func (ea *InMemoryEateryAdapter) FindEateriesByRank(rank int) ([]*entity.Eatery, *apperr.AppErr) {
	var res []*entity.Eatery

	for _, item := range eateries {
		if item.Rank == rank {
			res = append(res, item)
		}
	}

	return res, nil
}

func (ea *InMemoryEateryAdapter) FindEateriesByCategory(category string) ([]*entity.Eatery, *apperr.AppErr) {
	var res []*entity.Eatery

	for _, item := range eateries {
		if item.Category == category {
			res = append(res, item)
		}
	}

	return res, nil
}
