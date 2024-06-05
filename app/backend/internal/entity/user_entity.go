package entity

import (
	"github.com/google/uuid"
)

type UserEntity struct {
	ID       int64     `json:"id"`
	UUID     uuid.UUID `json:"uuid"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password,omitempty"`
}
