package response

import "github.com/foliveiracamara/bents-api/application/entity"

type UserResponse struct {
	UUID       string   `json:"uuid"`
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	Age        int      `json:"age"`
	Favorites  []string `json:"favorites"`
	FirstLogin bool     `json:"first_login"`
}

func (u *UserResponse) ParseUserDomainToResponse(user *entity.User) *UserResponse {
	return &UserResponse{
		UUID:       user.UUID,
		Name:       user.Name,
		Email:      user.Email,
		Age:        user.Age,
		Favorites:  user.Favorites,
		FirstLogin: user.FirstLogin,
	}
}