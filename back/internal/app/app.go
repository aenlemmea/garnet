package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aenlemmea/garnet/back/internal/api"
)

type AppContext struct {
	Logger              *log.Logger
	State               string
	AggregationHandler  *api.AggregationHandler
	PersonalizedHandler *api.PersonalizedHandler
	NewsHandler         *api.NewsHandler
}

func CreateAppContext() (*AppContext, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	aggregationHandler := api.CreateAggregationHandler()
	personalizedHandler := api.CreatePersonalizedHandler()
	newsHandler := api.CreateNewsHandler()

	appCtx := &AppContext{
		Logger:              logger,
		State:               "Running",
		AggregationHandler:  aggregationHandler,
		PersonalizedHandler: personalizedHandler,
		NewsHandler:         newsHandler,
	}

	return appCtx, nil
}

func (appCtx *AppContext) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "🟢 Status is available\n")
}
