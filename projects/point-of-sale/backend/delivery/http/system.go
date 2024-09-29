package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) registerSystemRoutes(r *mux.Router) {
	r.HandleFunc("/system/ping", a.handleHello)
}

func (a *API) handleHello(w http.ResponseWriter, r *http.Request) {
	jsonStringResponse(w, http.StatusOK, "OK!")
}
