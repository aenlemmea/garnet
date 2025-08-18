package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aenlemmea/garnet/back/internal/api"
	"github.com/aenlemmea/garnet/back/internal/data"
	"github.com/aenlemmea/garnet/back/internal/service/fetch"
	"github.com/aenlemmea/garnet/back/migrations"
)

type AppContext struct {
	Logger              *log.Logger
	State               string
	AggregationHandler  *api.AggregationHandler
	PersonalizedHandler *api.PersonalizedHandler
	NewsHandler         *api.NewsHandler
	FetcherService      fetch.NewsFetcher
	DB                  *sql.DB
}

func CreateAppContext() (*AppContext, error) {

	// Data Layer
	pgDB, err := data.Open()

	if err != nil {
		return nil, fmt.Errorf("db: Open error in Context. %w", err)
	}

	err = data.MigrateFS(pgDB, migrations.FS, ".")

	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	aggStore := data.CreatePostgresAggregatorStore(pgDB)
	perStore := data.CreatePostgresPersonalizedStore(pgDB)

	// API Layer

	aggregationHandler := api.CreateAggregationHandler(aggStore)
	personalizedHandler := api.CreatePersonalizedHandler(perStore)
	newsHandler := api.CreateNewsHandler(aggStore)

	// Service Layer

	fetchservice := fetch.CreateNewsAPIFetcher(aggStore, logger)

	appCtx := &AppContext{
		Logger:              logger,
		State:               "Running",
		AggregationHandler:  aggregationHandler,
		PersonalizedHandler: personalizedHandler,
		NewsHandler:         newsHandler,
		FetcherService:      fetchservice,
		DB:                  pgDB,
	}

	return appCtx, nil
}

func (appCtx *AppContext) RefreshFetch(w http.ResponseWriter, r *http.Request) {
	if err := appCtx.FetcherService.StartFetch(); err != nil {
		appCtx.Logger.Printf("Fetch error: %v\n", err)
		http.Error(w, "Failed to refresh data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Refresh successful")
}

func (appCtx *AppContext) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "🟢 Status is available\n")
}
