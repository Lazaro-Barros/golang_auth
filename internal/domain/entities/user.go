package entities

import (
	"github.com/google/uuid"

	"github.com/severusTI/auth_golang/pkg/ops"
	"github.com/severusTI/auth_golang/pkg/password"
	"github.com/severusTI/auth_golang/pkg/validations"
)

type User struct {
	id             string
	name           string
	email          string
	phoneNumber    string
	hashedPassword string
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
	encriptedPassword, err := password.HashPassword(inputedPassword)
	if err != nil {
		return nil, ops.Err(err)
	}

	return &User{
		id:             uuid.New().String(),
		email:          email,
		hashedPassword: encriptedPassword,
		name:           name,
		phoneNumber:    phoneNumber,
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

func (obj User) HashedPassword() string {
	return obj.hashedPassword
}

func (usr *User) FillFromData(id, name, email, phoneNumber, hashedPassword string) {
	usr.id = id
	usr.name = name
	usr.email = email
	usr.phoneNumber = phoneNumber
	usr.hashedPassword = hashedPassword
}
