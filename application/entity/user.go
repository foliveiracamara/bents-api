package entity

import (
	// "github.com/foliveiracamara/bents-api/configuration/apperr"
	// "golang.org/x/crypto/bcrypt"
)

type User struct {
	UUID       string
	Name       string
	Email      string
	Age        int
	Favorites  []string
	FirstLogin bool
	Password   string
}

// func (ud *User) EncryptPassword() *apperr.AppErr {
// 	hashedPassword, err := bcrypt.GenerateFromPassword(
// 		[]byte(ud.Password), 
// 		bcrypt.DefaultCost,
// 	)
// 	if err != nil {
// 		appErr := apperr.NewInternalServerError("Error encrypting password.")
// 		return appErr
// 	}
// 	ud.Password = string(hashedPassword)
// 	return nil
// }

// func (ud *User) CompareHashAndPassword(password string) *apperr.AppErr {
// 	err := bcrypt.CompareHashAndPassword([]byte(ud.Password), []byte(password))
// 	if err != nil {
// 		appErr := apperr.NewUnauthorizedError("Invalid password.")
// 		return appErr
// 	}
	
// 	return nil
// }
