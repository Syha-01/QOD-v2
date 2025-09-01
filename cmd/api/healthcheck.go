package main

import (
	"net/http"
)

const version = "1.0.0"

func (a *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": a.config.Environment,
			"version":     version,
		},
	}
	err := a.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}
