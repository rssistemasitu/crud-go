package model

import (
	"fmt"

	"github.com/rssistemasitu/crud-go/internal/configs/logger"
	"golang.org/x/crypto/bcrypt"
)

// Objeto que contém as regras de negócio
const (
	SECRET = "secret"
)

func (ud *userDomain) EncryptPassword() {
	passWithSecret := ud.GetPassword() + SECRET
	bytes, err := bcrypt.GenerateFromPassword([]byte(passWithSecret), bcrypt.DefaultCost)

	res, err := string(bytes), err

	if err != nil {
		logger.Error("Error to parse password", err)
	}

	fmt.Println("Hash: ", res)

	ud.password = res
}

func (ud *userDomain) CheckPasswordHash(pass string) bool {
	passWithSecret := pass + SECRET
	err := bcrypt.CompareHashAndPassword([]byte(ud.GetPassword()), []byte(passWithSecret))
	return err == nil
}
