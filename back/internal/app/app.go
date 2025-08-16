package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aenlemmea/garnet/back/internal/api"
	"github.com/aenlemmea/garnet/back/internal/data"
	"github.com/aenlemmea/garnet/back/migrations"
)

type AppContext struct {
	Logger              *log.Logger
	State               string
	AggregationHandler  *api.AggregationHandler
	PersonalizedHandler *api.PersonalizedHandler
	NewsHandler         *api.NewsHandler
	DB                  *sql.DB
}

func CreateAppContext() (*AppContext, error) {
	pgDB, err := data.Open()

	if err != nil {
		return nil, fmt.Errorf("db: Open error in Context. %w", err)
	}

	err = data.MigrateFS(pgDB, migrations.FS, ".")

	if err != nil {
		panic(err)
	}

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
		DB:                  pgDB,
	}

	return appCtx, nil
}

func (appCtx *AppContext) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "🟢 Status is available\n")
}
