package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
	"github.com/tijanadmi/tis-api/models"
)

func (app *application) getOneSignal(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	signalID, err := strconv.Atoi(id)
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	signals, err := app.models.DB.OneSignal(signalID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getDvDidf(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetDvDidf()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
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
	err = app.writeJSON(w, http.StatusOK, signals)
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
	err = app.writeJSON(w, http.StatusOK, signals)
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
	err = app.writeJSON(w, http.StatusOK, signals)
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
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getOneMalfunctionIn(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	signalID, err := strconv.Atoi(id)
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	signals, err := app.models.DB.OneMalfunctionIn(signalID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
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
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getOneAPU(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	signalID, err := strconv.Atoi(id)
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	signals, err := app.models.DB.OneAPU(signalID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
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
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getOCTR12(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetOCTR12()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getOCTRR(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetOCTRR()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getOCSP(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetOCSP()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getEarthfaultOCDV(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetEarthfaultOCDV()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getEarthfaultOCTR(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetEarthfaultOCTR()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getEarthfaultOCSP(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetEarthfaultOCSP()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getDirEarthfaultOC(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetDirEarthfaultOC()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getTPSendRcdv(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetTPSendRcdv()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getCircuitbreaker(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetCircuitbreaker()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getBBPBFtrip(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetBBPBFtrip()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getNonElectrical(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetNonElectrical()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getBBPBBtrip(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetBBPBBtrip()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getBFtrip(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetBFtrip()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getGroupsOfCauses(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetGroupsOfCauses()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getOneGroupOfCauses(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	signalID, err := strconv.Atoi(id)
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	signals, err := app.models.DB.OneGroupOfCauses(signalID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getCauses(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetCauses()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getOneCause(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	signalID, err := strconv.Atoi(id)
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	signals, err := app.models.DB.OneCause(signalID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getGroupOfReasons(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetGroupOfReasons()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getOneGroupOfReason(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	signalID, err := strconv.Atoi(id)
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	signals, err := app.models.DB.OneGroupOfReasons(signalID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getReasons(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetReasons()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getOneReason(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	signalID, err := strconv.Atoi(id)
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	signals, err := app.models.DB.OneReason(signalID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getWeatherConditions(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetWeatherConditions()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getOneWeatherCondition(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	signalID, err := strconv.Atoi(id)
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	signals, err := app.models.DB.OneWeatherCondition(signalID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getCategoriesOfEvents(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetCategoriesOfEvents()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getOneCategoryOfEvents(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	signalID, err := strconv.Atoi(id)
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	signals, err := app.models.DB.OneCategoryOfEvents(signalID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getOHL(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetOHL()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getPowerCables(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetPowerCables()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getSubstations(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetSubstations()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getFeeders(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetFeeders()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getProtectionDevices(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetProtectionDevices()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getPowerTransformers(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetPowerTransformers()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getDisconnectors(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetDisconnectors()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getWorkPermissions(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetWorkPermissions()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getWorkInEENetwork(w http.ResponseWriter, r *http.Request) {
	signals, err := app.models.DB.GetWorkInEENetwork()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, signals)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getWeather(w http.ResponseWriter, r *http.Request) {
	year := chi.URLParam(r, "year")
	weatherData, err := app.models.DB.GetWeather(year)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, weatherData)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getWeatherForecast(w http.ResponseWriter, r *http.Request) {

	weatherData, err := app.models.DB.GetWeatherForecast()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, weatherData)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getWeatherHistory(w http.ResponseWriter, r *http.Request) {
	year := chi.URLParam(r, "year")

	weatherData, err := app.models.DB.GetWeatherHistory(year)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, weatherData)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getPermissions1(w http.ResponseWriter, r *http.Request) {
	year := chi.URLParam(r, "year")

	weatherData, err := app.models.DB.GetPermissions1(year)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, weatherData)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getPermissions23(w http.ResponseWriter, r *http.Request) {
	year := chi.URLParam(r, "year")

	weatherData, err := app.models.DB.GetPermissions23(year)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, weatherData)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getRequests1(w http.ResponseWriter, r *http.Request) {
	year := chi.URLParam(r, "year")

	weatherData, err := app.models.DB.GetRequests1(year)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, weatherData)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getRequests23(w http.ResponseWriter, r *http.Request) {
	year := chi.URLParam(r, "year")

	weatherData, err := app.models.DB.GetRequests23(year)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, weatherData)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getOutages(w http.ResponseWriter, r *http.Request) {
	year := chi.URLParam(r, "year")

	weatherData, err := app.models.DB.GetOutages(year)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, weatherData)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getExclusions(w http.ResponseWriter, r *http.Request) {
	year := chi.URLParam(r, "year")

	weatherData, err := app.models.DB.GetExclusions(year)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, weatherData)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getPlans(w http.ResponseWriter, r *http.Request) {
	year := chi.URLParam(r, "year")

	weatherData, err := app.models.DB.GetPlans(year)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, weatherData)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

/***** start PiPiDDN api  ****/

func (app *application) insertPiPiDDNIsklj(w http.ResponseWriter, r *http.Request) {
	var payload models.PiPiDDNIsklj

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {

		app.errorJSON(w, err)
		return
	}
	fmt.Printf("Izgleda da je TipMan %s\n", payload.TipMan)
	err = app.models.DB.InsertPiPiDDNIsklj(payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, "Poruka message")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) insertPiPiDDNIspad(w http.ResponseWriter, r *http.Request) {
	var payload models.PiPiDDNIspad

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {

		app.errorJSON(w, err)
		return
	}
	//fmt.Printf("Izgleda da je TipMan %s\n", payload.TipMan)
	err = app.models.DB.InsertPiPiDDNIspad(payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, "Poruka message")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) updatePiPiDDNIsklj(w http.ResponseWriter, r *http.Request) {
	var payload models.PiPiDDNIsklj

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.models.DB.UpdatePiPiDDNIsklj(payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponse{
		Error:   false,
		Message: "pipiddn updated",
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}

func (app *application) updatePiPiDDNIspad(w http.ResponseWriter, r *http.Request) {
	var payload models.PiPiDDNIspad

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.models.DB.UpdatePiPiDDNIspad(payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponse{
		Error:   false,
		Message: "pipiddn updated",
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}

func (app *application) deletePiPiDDN(w http.ResponseWriter, r *http.Request) {
	//id, err := strconv.Atoi(chi.URLParam(r, "id"))
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.models.DB.DeletePiPiDDN(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponse{
		Error:   false,
		Message: "pipiddn deleted",
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}

/*** end PiPiDDN api  ***/

/**************** the another method ***********/
func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	// read json payload
	var requestPayload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate user against database
	user, err := app.models.DB.GetUserByUsername(requestPayload.Username)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// check password
	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	var roles []string
	for _, r := range user.UserRole {
		roles = append(roles, r.RoleCode)
	}

	//fmt.Println(roles)
	// create a jwt user
	u := jwtUser{
		ID:        user.ID,
		FirstName: user.Username,
		LastName:  user.Username,
		Roles:     roles,
	}

	// generate tokens
	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	//log.Println(tokens.Token)

	refreshCookie := app.auth.GetRefreshCookie(tokens.RefreshToken)
	http.SetCookie(w, refreshCookie)

	app.writeJSON(w, http.StatusAccepted, tokens)
}

func (app *application) refreshToken(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		if cookie.Name == app.auth.CookieName {
			claims := &Claims{}
			refreshToken := cookie.Value

			// parse the token to get the claims
			_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(app.JWTSecret), nil
			})
			if err != nil {
				app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
				return
			}

			// get the user id from the token claims
			userID, err := strconv.Atoi(claims.Subject)
			if err != nil {
				app.errorJSON(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			user, err := app.models.DB.GetUserByID(userID)
			if err != nil {
				app.errorJSON(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			u := jwtUser{
				ID:        user.ID,
				FirstName: user.Username,
				LastName:  user.Username,
			}

			tokenPairs, err := app.auth.GenerateTokenPair(&u)
			if err != nil {
				app.errorJSON(w, errors.New("error generating tokens"), http.StatusUnauthorized)
				return
			}

			http.SetCookie(w, app.auth.GetRefreshCookie(tokenPairs.RefreshToken))

			app.writeJSON(w, http.StatusOK, tokenPairs)

		}
	}
}

func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, app.auth.GetExpiredRefreshCookie())
	w.WriteHeader(http.StatusAccepted)
}
