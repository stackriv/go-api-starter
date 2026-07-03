package interfaces

import (
	"github.com/google/uuid"
	"github.com/stackriv/go-api-starter/internal/business/model"
)

type UserRepository interface {
	FindByID(id uuid.UUID) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Save(user *model.User) error
	CountUsers() (uint, error)
}
