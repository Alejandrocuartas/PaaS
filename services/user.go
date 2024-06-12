package services

import (
	"PaaS/models"
	"PaaS/repositories"
	"PaaS/utilities"

	"errors"

	"github.com/jinzhu/gorm"
)

func Signup(data models.Signup) (
	r models.SignupResponse,
	e error,
) {

	hash, e := utilities.HashPassword(data.Password)
	if e != nil {
		return r, utilities.ManageError(e)
	}

	user := models.User{
		Email:    data.Email,
		Name:     data.Name,
		Password: hash,
	}

	e = repositories.CreateUser(&user)
	if e != nil {
		return r, utilities.ManageError(e)
	}

	r = models.SignupResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return r, e
}

func Login(data models.Login) (
	r models.LoginResponse,
	e error,
) {
	user, e := repositories.GetUserByEmail(data.Email)
	if e != nil {
		if errors.Is(e, gorm.ErrRecordNotFound) {
			return r, errors.New("user with this email does not exist")
		}
		return r, utilities.ManageError(e)
	}

	if !utilities.ComparePassword(user.Password, data.Password) {
		return r, errors.New("invalid password")
	}

	r = models.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return r, e
}
