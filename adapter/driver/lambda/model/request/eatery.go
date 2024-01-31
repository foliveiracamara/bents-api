package request

type EateryLoginRequest struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=8"`
}

type EateryRequest struct {
	UUID     string `json:"uuid,omitempty"`
	Name     string `json:"name,omitempty" validate:"required,min=3,max=60"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Category string `json:"Category,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required,min=8,max=16"`

	// Remove later, doesnt make sense in eatery creation
	Rank int `json:"rank,omitempty"`
}
