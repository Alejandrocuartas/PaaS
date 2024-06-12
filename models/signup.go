package models

import (
	"errors"
)

type Signup struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (s Signup) Validate() error {
	if s.Email == "" {
		return errors.New("email is required")
	}
	if s.Password == "" {
		return errors.New("password is required")
	}
	if s.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

type SignupResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
