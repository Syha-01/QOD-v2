package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// routes returns a new http.Handler with the application's routes.
func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	return router
}
