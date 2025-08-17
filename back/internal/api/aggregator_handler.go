package api

import (
	"fmt"
	"net/http"

	"github.com/aenlemmea/garnet/back/internal/data"
	"github.com/go-chi/chi/v5"
)

type AggregationHandler struct {
	aggregatorStore data.AggregatorStore
}

func CreateAggregationHandler(aggregatorStore data.AggregatorStore) *AggregationHandler {
	return &AggregationHandler{
		aggregatorStore: aggregatorStore,
	}
}

func (ah *AggregationHandler) HandleGetAggregation(w http.ResponseWriter, r *http.Request) {

}

func (ah *AggregationHandler) HandlePostRefresh(w http.ResponseWriter, r *http.Request) {
	paramsUserID := chi.URLParam(r, "uid")

	if paramsUserID == "" {
		http.Error(w, "Not valid account", http.StatusForbidden)
		return
	}
	fmt.Fprintf(w, "This is the user ID: %s\n", paramsUserID)
}

func (ah *AggregationHandler) HandleGetCategory(w http.ResponseWriter, r *http.Request) {}
