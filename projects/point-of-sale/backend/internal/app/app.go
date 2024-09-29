package app

import (
	"backend/internal/store"
	"backend/model"
	"context"
)

type app struct {
	store store.Store
}

type App interface {
	GetServerMetadata() *model.ServerMetadata

	// Auth
	Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error)
	Register(ctx context.Context, request *model.CreateUserRequest) error
}

func New(store store.Store) App {
	app := &app{
		store: store,
	}

	return app
}
