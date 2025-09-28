package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Email        Email
	Name         Name
	PasswordHash PasswordHash
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
type Email struct {
	Value string
}

type Name struct {
	First string
	Last  string
}

type PasswordHash struct {
	Value string
}
