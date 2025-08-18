package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aenlemmea/garnet/back/internal/data"
	"github.com/go-chi/chi/v5"
)

type NewsHandler struct {
	aggregatorStore data.AggregatorStore
} // Only knows about aggregator store.

func CreateNewsHandler(aggregatorStore data.AggregatorStore) *NewsHandler {
	return &NewsHandler{
		aggregatorStore: aggregatorStore,
	}
}

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

	agg, err := ph.aggregatorStore.GetAggergatorByID(newsID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed reading news by id", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(agg)
}
