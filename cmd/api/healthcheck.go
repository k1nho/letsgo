package main

import (
	"net/http"
)

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
		app.logger.Print(err)
		http.Error(w, "The server could not process request", http.StatusInternalServerError)
		return
	}

}
