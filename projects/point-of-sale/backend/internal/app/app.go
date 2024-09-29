package app

import "backend/internal/store"

type App struct {
	store store.Store
}

type Services struct {
	Store store.Store
}

func New(services Services) *App {
	app := &App{
		store: services.Store,
	}

	return app
}
