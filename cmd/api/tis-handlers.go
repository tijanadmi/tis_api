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
	err = app.writeJSON(w, http.StatusOK, signals, "dv_didff_signals")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getDiffTr(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetDiffTr()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals, "tr_diff_signals")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getDisTrRes(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetDisTrRes()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals, "tr_dis_res_signals")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getDisDiffSp(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetDisDiffSp()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals, "sp_dis_diff_signals")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getMalfunctionIn(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetMalfunctionIn()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals, "malfunction_in")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAPU(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetAPU()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals, "apu")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getOCDV(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetOCDV()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals, "dv_oc_protection")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}
