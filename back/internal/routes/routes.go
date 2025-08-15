package routes

import (
	"github.com/aenlemmea/garnet/back/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(appCtx *app.AppContext) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", appCtx.HealthCheck)

	return r
}
