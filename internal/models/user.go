package models

import (
	"errors"
	"marketplace/pkg/validation"
)

type User struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Validate() error {
	errLog := validation.ValidLogin(u.Login)
	errPas := validation.ValidPassword(u.Password)
	err := ""
	if errLog != nil {
		err += errLog.Error() + "\n"
	}
	if errPas != nil {
		err += errPas.Error()
	}
	if len(err) != 0 {
		return errors.New(err)
	}
	return nil
}
