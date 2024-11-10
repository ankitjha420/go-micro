package main

import (
	"logger-service/data"
	"net/http"
)

type jsonPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var requestPayload jsonPayload
	_ = app.readJSON(w, r, &requestPayload)

	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		_ = app.errorJSON(w, err)
		return
	}

	response := jsonResponse{
		Error:   false,
		Message: "logged",
	}

	_ = app.writeJSON(w, http.StatusAccepted, response)
}