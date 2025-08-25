package main

import (
	"net/http"
)

const version = "1.0.0"

// healthcheckHandler returns the health of the system.
func (a *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"status": "available", "environment": a.config.env, "version": version}

	err := a.writeJSON(w, http.StatusOK, data, nil)

	if err != nil {
		a.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
