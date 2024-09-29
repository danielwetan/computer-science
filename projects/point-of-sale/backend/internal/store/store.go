package store

import (
	"backend/model"
	"context"
)

type Store interface {
	// General
	DBType() string
	DBVersion() string

	// Users
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, request *model.CreateUserRequest) error
}