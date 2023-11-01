package view

import (
	"github.com/rssistemasitu/crud-go/internal/controller/model/response"
	"github.com/rssistemasitu/crud-go/internal/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetId(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   int8(userDomain.GetAge()),
	}

}
