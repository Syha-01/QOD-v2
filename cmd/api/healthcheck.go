package main

import (
	"fmt"
	"net/http"
)

const version = "1.0.0"

// healthcheckHandler returns the health of the system.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	js := `{"status": "available", "environment": %q, "version": %q}`
	js = fmt.Sprintf(js, app.config.env, version)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(js))
}
