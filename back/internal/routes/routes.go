package routes

import (
	"github.com/aenlemmea/garnet/back/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(appCtx *app.AppContext) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", appCtx.HealthCheck)

	r.Get("/news/{id}", appCtx.NewsHandler.HandleGetNewsById)

	r.Get("/aggregation/", appCtx.AggregationHandler.HandleGetAggregation)
	r.Get("/aggregation/categories", appCtx.AggregationHandler.HandleGetCategory)
	r.Post("/aggregation/refresh/{uid}", appCtx.AggregationHandler.HandlePostRefresh)

	r.Get("/personalized/{uid}", appCtx.PersonalizedHandler.HandleGetPersonalized)
	r.Post("/personalized/refresh/{uid}", appCtx.PersonalizedHandler.HandlePostRefresh)
	return r
}
