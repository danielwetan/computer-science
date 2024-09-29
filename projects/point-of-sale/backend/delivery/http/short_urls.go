package api

import (
	"backend/model"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) registerShortUrlRoutes(r *mux.Router) {
	// r.HandleFunc("/short_urls", a.secureRoutes(a.handleCreateShortUrl)).Methods(http.MethodPost)
	r.HandleFunc("/short_urls/public", a.handleCreateShortUrlPublic).Methods(http.MethodPost)
	r.HandleFunc("/short_urls/{code}", a.handleGetShortUrl).Methods(http.MethodGet)
}

func (a *API) handleCreateShortUrlPublic(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		errorResponse(w, err)
		return
	}

	var createShortUrlRequest model.CreateShortUrlRequest
	err = json.Unmarshal(requestBody, &createShortUrlRequest)
	if err != nil {
		errorResponse(w, err)
		return
	}

	createShortUrlResponse, err := a.app.CreateShortUrl(r.Context(), &createShortUrlRequest)
	if err != nil {
		errorResponse(w, err)
		return
	}

	response, err := json.Marshal(createShortUrlResponse)
	if err != nil {
		errorResponse(w, err)
		return
	}

	jsonStringResponse(w, http.StatusOK, string(response))
}

func (a *API) handleGetShortUrl(w http.ResponseWriter, r *http.Request) {
	request := &model.GetShortUrlRequest{
		ShortCode: mux.Vars(r)["code"],
	}

	getShortUrlResponse, err := a.app.GetShortUrl(r.Context(), request)
	if err != nil {
		errorResponse(w, err)
		return
	}

	response, err := json.Marshal(getShortUrlResponse)
	if err != nil {
		errorResponse(w, err)
		return
	}

	jsonStringResponse(w, http.StatusOK, string(response))
}
