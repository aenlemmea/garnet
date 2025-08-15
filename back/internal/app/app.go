package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type AppContext struct {
	Logger *log.Logger
	State  string
}

func CreateAppContext() (*AppContext, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	appCtx := &AppContext{
		Logger: logger,
		State:  "Running",
	}

	return appCtx, nil
}

func (appCtx *AppContext) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "🟢 Status is available\n")
}
