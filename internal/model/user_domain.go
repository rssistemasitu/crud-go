package model

import (
	"encoding/json"

	"github.com/rssistemasitu/crud-go/internal/configs/logger"
)

type userDomain struct {
	id       string
	email    string
	password string
	name     string
	age      int
}

func (ud *userDomain) GetId() string {
	return ud.id
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetAge() int {
	return ud.age
}

func (ud *userDomain) SetId(id string) {
	ud.id = id
}

func (ud *userDomain) GetJsonValue() (string, error) {
	b, err := json.Marshal(ud)
	if err != nil {
		logger.Error("Error to parse domain", err)
		return "", err
	}
	return string(b), nil
}
