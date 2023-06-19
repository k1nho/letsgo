package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.WriteJson(w, http.StatusOK, data, nil)

	if err != nil {
		app.logger.Print(err)
		http.Error(w, "The server could not process request", http.StatusInternalServerError)
		return
	}

}
