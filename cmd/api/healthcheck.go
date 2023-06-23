package main

import (
	"net/http"
)

// healthcheckHandler: PING endpoint to check status
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	envelope := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.WriteJson(w, http.StatusOK, envelope, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
