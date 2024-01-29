package security

import (
	"github.com/foliveiracamara/bents-api/configuration/apperr"
	"golang.org/x/crypto/bcrypt"
)

type Encryption struct {}

func (e *Encryption) EncryptPassword(password string) (string, *apperr.AppErr) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		appErr := apperr.NewInternalServerError("Error encrypting password.")
		return "", appErr
	}
	password = string(hashedPassword)
	
	return password, nil
}

func (e *Encryption) CompareHashAndPassword(hashedPassword string, password string) *apperr.AppErr {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		appErr := apperr.NewUnauthorizedError("Invalid password.")
		return appErr
	}
	
	return nil
}