package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type NewsHandler struct{} // Only knows about aggregator store.

func (ph *NewsHandler) HandleGetNewsById(w http.ResponseWriter, r *http.Request) {
	paramsNewsID := chi.URLParam(r, "id")
	if paramsNewsID == "" {
		http.NotFound(w, r)
		return
	}

	newsID, err := strconv.ParseInt(paramsNewsID, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "This is the news ID: %d\n", newsID)
}
