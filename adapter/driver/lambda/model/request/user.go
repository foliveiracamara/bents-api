package request

type UserLoginRequest struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=8"`
}

type UserRequest struct {
	UUID       string   `json:"uuid,omitempty"`
	Name       string   `json:"name,omitempty" validate:"required,min=3,max=40"`
	LastName   string   `json:"last_name,omitempty" validate:"required,min=3,max=120"`
	Email      string   `json:"email,omitempty" validate:"required,email"`
	Age        int      `json:"age,omitempty" validate:"required,min=1,max=110"`
	Favorites  []string `json:"favorites,omitempty"`
	FirstLogin bool     `json:"first_login,omitempty"`
	Password   string   `json:"password,omitempty" validate:"required,min=8,max=16"`
}
