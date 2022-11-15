package main

import (
	"net/http"
)

func (app *application) getDvDidf(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetDvDidf()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals, "dv_didf_signals")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}
