package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type PersonalizedHandler struct{}

func CreatePersonalizedHandler() *PersonalizedHandler {
	return &PersonalizedHandler{}
}

func (ph *PersonalizedHandler) HandlePostRefresh(w http.ResponseWriter, r *http.Request) {
	paramsUserID := chi.URLParam(r, "uid")

	if paramsUserID == "" {
		http.Error(w, "Not valid account", http.StatusForbidden)
		return
	}
	fmt.Fprintf(w, "This is the user ID: %s\n", paramsUserID)
}

func (ph *PersonalizedHandler) HandleGetPersonalized(w http.ResponseWriter, r *http.Request) {
	paramsUserID := chi.URLParam(r, "uid")

	if paramsUserID == "" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "This is the user ID: %s\n", paramsUserID)
}
