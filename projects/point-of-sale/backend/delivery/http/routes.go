package api

import "github.com/gorilla/mux"

func (a *API) RegisterRoutes(r *mux.Router) {
	v1 := r.PathPrefix("/v1").Subrouter()
	a.registerSystemRoutes(v1)
	a.registerAuthRoutes(v1)
}
