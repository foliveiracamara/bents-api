package response

import "github.com/foliveiracamara/bents-api/application/entity"

type EateryResponse struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Type  string `json:"type"`
}

type EateriesResponse struct {
	Eateries []*EateryResponse `json:"eateries"`
}

func (u *EateryResponse) ParseEateryDomainToResponse(eat *entity.Eatery) *EateryResponse {
	return &EateryResponse{
		UUID:  eat.UUID,
		Name:  eat.Name,
		Email: eat.Email,
		Type:  eat.Type,
	}
}

// func (u *EateriesResponse) ParseEateryDomainToResponse(eat []*entity.Eatery) *EateriesResponse {
// 	for _, e := range eat {
// 		return &EateriesResponse{
// 			Eateries: []*EateryResponse{
// 				UUID: e.UUID,
// 				Name: e.Name,
// 				Email: e.Email,
// 				Type: e.Type,
// 			},
// 		},
// 	}
// 	return 
// }

