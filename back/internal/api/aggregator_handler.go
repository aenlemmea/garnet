package api

import "net/http"

type AggregationHandler struct{}

func CreateAggregationHandler() *AggregationHandler {
	return &AggregationHandler{}
}

func (ah *AggregationHandler) HandleGetAggregation(w http.ResponseWriter, r *http.Request) {

}

func (ah *AggregationHandler) HandlePostRefresh(w http.ResponseWriter, r *http.Request) {

}

func (ah *AggregationHandler) HandleGetCategory(w http.ResponseWriter, r *http.Request) {}
