package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/aenlemmea/garnet/back/internal/app"
	"github.com/aenlemmea/garnet/back/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8090, "The port to run backend app on.")
	flag.Parse()
	appCtx, err := app.CreateAppContext()
	if err != nil {
		panic(err)
	}

	r := routes.SetupRoutes(appCtx)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	appCtx.Logger.Printf("%s on port: %d\n", appCtx.State, port)

	err = server.ListenAndServe()

	if err != nil {
		appCtx.Logger.Fatal(err)
	}
}
