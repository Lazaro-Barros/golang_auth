package domain

import (
	"github.com/google/uuid"
	"github.com/severusTI/auth_golang/pkg/ops"
	"github.com/severusTI/auth_golang/pkg/password"
	"github.com/severusTI/auth_golang/pkg/validations"
)

type User struct {
	id                string `json:""`
	name              string `json:""`
	email             string `json:""`
	phoneNumber       string `json:""`
	encriptedPassword string `json:""`
}

func NewUser(email, inputedPassword, name, phoneNumber string) (user *User, err error) {

	if err = validations.IsValidEmail(email); err != nil {
		return nil, ops.Err(err)
	}

	if err = validations.IsValidPassword(inputedPassword); err != nil {
		return nil, ops.Err(err)
	}

	if err = validations.IsValidPhoneNumber(phoneNumber); err != nil {
		return nil, ops.Err(err)
	}
	encriptedPassword, err := password.EncryptPassword(inputedPassword)
	if err != nil {
		return nil, ops.Err(err)
	}

	return &User{
		id:                uuid.New().String(),
		email:             email,
		encriptedPassword: encriptedPassword,
		name:              name,
		phoneNumber:       phoneNumber,
	}, nil
}

func (usr *User) Id() string {
	return usr.id
}

func (usr *User) Name() string {
	return usr.name
}

func (usr *User) Email() string {
	return usr.email
}

func (usr *User) PhoneNumber() string {
	return usr.phoneNumber
}

func (obj User) EncriptedPassword() string {
	return obj.encriptedPassword
}

func (usr *User) FillFromData(id, name, email, phoneNumber, encriptedPassword string) {
	usr.id = id
	usr.name = name
	usr.email = email
	usr.phoneNumber = phoneNumber
	usr.encriptedPassword = encriptedPassword
}
