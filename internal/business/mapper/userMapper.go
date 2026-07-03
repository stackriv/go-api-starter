package mapper

import (
	"github.com/stackriv/go-api-starter/internal/business/dto"
	"github.com/stackriv/go-api-starter/internal/business/model"
)

func ToRequest(user *model.User) *dto.UserRequest {
	return &dto.UserRequest{
		Email:    user.Email,
		Password: user.Password,
	}
}

func toUser(request *dto.UserRequest) *model.User {
	return &model.User{
		Email:    request.Email,
		Password: request.Password,
	}
}

func toResponse(user *model.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:    user.ID,
		Email: user.Email,
	}
}
