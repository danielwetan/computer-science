package api

import (
	"backend/model"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) registerShortUrlRoutes(r *mux.Router) {
	// r.HandleFunc("/short_urls", a.secureRoutes(a.handleCreateShortUrl)).Methods(http.MethodPost)
	r.HandleFunc("/short_urls/public", a.handleCreateShortUrlPublic).Methods(http.MethodPost)

	// simplify the redirection logic
	// for the real-world app we can separate this as a different service with custom domain
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

	if getShortUrlResponse.Target == "" {
		errorResponse(w, errors.New("invalid short code or URL not found"))
		return
	}

	http.Redirect(w, r, getShortUrlResponse.Target, http.StatusFound)
}
