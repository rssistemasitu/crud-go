package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rssistemasitu/crud-go/internal/configs/logger"
	"github.com/rssistemasitu/crud-go/internal/rest_err"
	"github.com/rssistemasitu/crud-go/internal/utils"
)

const (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) {
	secret := utils.GetEnvVariable(JWT_SECRET_KEY)
	expirationTime := time.Now().Add(1 * time.Minute)

	claims := jwt.MapClaims{
		"id":    ud.id,
		"email": ud.email,
		"name":  ud.name,
		"age":   ud.age,
		"exp":   expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", rest_err.NewInternalServerError(fmt.Sprintf("Error trying to generate jwt token, err=%s", err.Error()))
	}

	logger.Info(fmt.Sprintf("Bearer: %s", tokenString))

	return tokenString, nil
}

func removeBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}
	return token
}

func VerifyToken(tokenString string) (UserDomainInterface, *rest_err.RestErr) {
	secret := utils.GetEnvVariable(JWT_SECRET_KEY)

	parsedToken, err := jwt.Parse(removeBearerPrefix(tokenString), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, rest_err.NewBadRequestError("Invalid token")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, rest_err.NewUnauthorizedError("Invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	expirationTime := int64(parsedToken.Claims.(jwt.MapClaims)["exp"].(float64))
	if time.Now().Unix() > expirationTime {
		return nil, rest_err.NewUnauthorizedError("Expired token")
	}

	if !ok || !parsedToken.Valid {
		return nil, rest_err.NewUnauthorizedError("Invalid token")
	}

	return &userDomain{
		id:    claims["id"].(string),
		email: claims["email"].(string),
		name:  claims["name"].(string),
		age:   int(claims["age"].(float64)),
	}, nil
}

func MiddlewareVerifyToken(c *gin.Context) {
	secret := utils.GetEnvVariable(JWT_SECRET_KEY)
	tokenValue := removeBearerPrefix(c.Request.Header.Get("Authorization"))

	parsedToken, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, rest_err.NewBadRequestError("Invalid token")
	})

	errRest := rest_err.NewUnauthorizedError("Invalid token")

	if err != nil {
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	_, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok || !parsedToken.Valid {
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

}
