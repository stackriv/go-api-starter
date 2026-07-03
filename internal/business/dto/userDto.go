package dto

import "github.com/google/uuid"

type UserRequest struct {
	Email    string
	Password string
}

type UserResponse struct {
	ID    uuid.UUID
	Email string
}
