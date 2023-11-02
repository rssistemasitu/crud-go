package model

import "github.com/rssistemasitu/crud-go/internal/rest_err"

type UserDomainInterface interface {
	GetId() string
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int
	SetId(string)
	GetJsonValue() (string, error)
	EncryptPassword()
	CheckPasswordHash(string) bool
	GenerateToken() (string, *rest_err.RestErr)
}

func NewUserDomain(
	email,
	password,
	name string,
	age int,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func NewUserUpdateDomain(
	name string,
	age int,
) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}
}

func NewUserLoginDomain(
	email, password string,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}
