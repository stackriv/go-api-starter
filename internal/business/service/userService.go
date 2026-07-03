package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/stackriv/go-api-starter/internal/business/dto"
	"github.com/stackriv/go-api-starter/internal/business/mapper"
	"github.com/stackriv/go-api-starter/internal/business/reposiroty/interfaces"
)

type UserService struct {
	Repository interfaces.UserRepository
}

func (u *UserService) GetUserByID(id uuid.UUID) (*dto.UserRequest, error) {
	user, err := u.Repository.FindByID(id)
	if err != nil {
		return nil, errors.New("failed to find user by ID")
	}
	return mapper.ToRequest(user), nil
}
