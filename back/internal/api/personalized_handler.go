package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type PersonalizedHandler struct{}

func (ph *PersonalizedHandler) HandlePostRefresh(w http.ResponseWriter, r *http.Request) {
	paramsUserID := chi.URLParam(r, "uid")

	if paramsUserID == "" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "This is the user ID: %s\n", paramsUserID)
}
