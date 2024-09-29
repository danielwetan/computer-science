package app

import (
	"backend/internal/store"
	"backend/model"
)

type app struct {
	store store.Store
}

type App interface {
	GetServerMetadata() *model.ServerMetadata
}

func New(store store.Store) App {
	app := &app{
		store: store,
	}

	return app
}
