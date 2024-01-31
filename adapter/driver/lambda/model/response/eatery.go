package response

import (
	"fmt"

	"github.com/foliveiracamara/bents-api/application/entity"
)

type EateryResponse struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Category string `json:"category"`
	Rank     int    `json:"rank"`
}

type EateriesResponse struct {
	Eateries []*EateryResponse `json:"eateries"`
}

func (u *EateryResponse) ParseEateryDomainToResponse(eat *entity.Eatery) *EateryResponse {
	return &EateryResponse{
		UUID:     eat.UUID,
		Name:     eat.Name,
		Email:    eat.Email,
		Category: eat.Category,
	}
}

func (u *EateriesResponse) ParseEateryDomainToResponse(eat []*entity.Eatery) *EateriesResponse {
	var eateries []*EateryResponse
	for _, e := range eat {
		fmt.Println(e)
		eateries = append(eateries, &EateryResponse{
			UUID:  e.UUID,
			Name:  e.Name,
			Email: e.Email,
			Rank:  e.Rank,	
			Category: e.Category,
		})
	}

	return &EateriesResponse{
		Eateries: eateries,
	}
}
