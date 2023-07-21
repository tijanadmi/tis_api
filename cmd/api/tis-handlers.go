package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"time"

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

	signals, err := app.DB.OneSignal(signalID)
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
	signals, err := app.DB.GetDvDidf()
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
	signals, err := app.DB.GetDiffTr()
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
	signals, err := app.DB.GetDisTrRes()
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
	signals, err := app.DB.GetDisDiffSp()
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
	signals, err := app.DB.GetMalfunctionIn()
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

	signals, err := app.DB.OneMalfunctionIn(signalID)
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
	signals, err := app.DB.GetAPU()
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

	signals, err := app.DB.OneAPU(signalID)
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
	signals, err := app.DB.GetOCDV()
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
	signals, err := app.DB.GetOCTR12()
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
	signals, err := app.DB.GetOCTRR()
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
	signals, err := app.DB.GetOCSP()
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
	signals, err := app.DB.GetEarthfaultOCDV()
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
	signals, err := app.DB.GetEarthfaultOCTR()
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
	signals, err := app.DB.GetEarthfaultOCSP()
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

func (app *application) getEarthfaultOCTRR(w http.ResponseWriter, r *http.Request) {
	signals, err := app.DB.GetEarthfaultOCTRR()
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
	signals, err := app.DB.GetDirEarthfaultOC()
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
	signals, err := app.DB.GetTPSendRcdv()
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
	signals, err := app.DB.GetCircuitbreaker()
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
	signals, err := app.DB.GetBBPBFtrip()
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
	signals, err := app.DB.GetNonElectrical()
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
	signals, err := app.DB.GetBBPBBtrip()
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
	signals, err := app.DB.GetBFtrip()
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
	signals, err := app.DB.GetGroupsOfCauses()
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

	signals, err := app.DB.OneGroupOfCauses(signalID)
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
	signals, err := app.DB.GetCauses()
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

	signals, err := app.DB.OneCause(signalID)
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
	signals, err := app.DB.GetGroupOfReasons()
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

	signals, err := app.DB.OneGroupOfReasons(signalID)
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
	signals, err := app.DB.GetReasons()
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

	signals, err := app.DB.OneReason(signalID)
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
	signals, err := app.DB.GetWeatherConditions()
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

	signals, err := app.DB.OneWeatherCondition(signalID)
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
	signals, err := app.DB.GetCategoriesOfEvents()
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

	signals, err := app.DB.OneCategoryOfEvents(signalID)
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
	signals, err := app.DB.GetOHL()
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
	signals, err := app.DB.GetPowerCables()
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
	signals, err := app.DB.GetSubstations()
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
	signals, err := app.DB.GetFeeders()
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
	signals, err := app.DB.GetProtectionDevices()
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
	signals, err := app.DB.GetPowerTransformers()
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
	signals, err := app.DB.GetDisconnectors()
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
	signals, err := app.DB.GetWorkPermissions()
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

func (app *application) getWorkPermissionsAll(w http.ResponseWriter, r *http.Request) {
	signals, err := app.DB.GetWorkPermissionsAll()
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

func (app *application) getRequest1Gr(w http.ResponseWriter, r *http.Request) {
	signals, err := app.DB.GetRequest1Gr()
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

func (app *application) getRequest2Gr(w http.ResponseWriter, r *http.Request) {
	signals, err := app.DB.GetRequest2Gr()
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
	signals, err := app.DB.GetWorkInEENetwork()
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
	weatherData, err := app.DB.GetWeather(year)
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

	weatherData, err := app.DB.GetWeatherForecast()
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

	weatherData, err := app.DB.GetWeatherHistory(year)
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

	weatherData, err := app.DB.GetPermissions1(year)
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

	weatherData, err := app.DB.GetPermissions23(year)
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

	weatherData, err := app.DB.GetRequests1(year)
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

	weatherData, err := app.DB.GetRequests23(year)
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

	weatherData, err := app.DB.GetOutages(year)
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

	weatherData, err := app.DB.GetExclusions(year)
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

	weatherData, err := app.DB.GetPlans(year)
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

func (app *application) getUnopenedPermitForDay(w http.ResponseWriter, r *http.Request) {
	day := chi.URLParam(r, "day")
	org := chi.URLParam(r, "org")

	unopenedPermits, err := app.DB.GetUnopenedPermitForDay(day, org)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, unopenedPermits)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

/***** start PiPiDDN api  ****/

func (app *application) getAllPiPiDDN(w http.ResponseWriter, r *http.Request) {
	//year := chi.URLParam(r, "year")

	p, err := app.DB.GetAllPiPiDDN()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, p)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllPiPiDDNP(w http.ResponseWriter, r *http.Request) {
	//year := chi.URLParam(r, "year")

	p, err := app.DB.GetAllPiPiDDNP()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, p)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getPiPiDDNByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	/*pipiddnID, err := strconv.Atoi(id)
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}*/

	p, err := app.DB.GetPiPiDDNByID(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, p)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getPiPiDDNByIDP(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	/*pipiddnID, err := strconv.Atoi(id)
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}*/

	p, err := app.DB.GetPiPiDDNByIDP(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, p)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) insertPiPiDDNIsklj(w http.ResponseWriter, r *http.Request) {
	var payload models.PiPiDDNIsklj

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {

		app.errorJSON(w, err)
		return
	}

	var resp JSONResponse
	if payload.SynsoftId == "" || payload.DatSmene == "" || payload.IdSGrraz == "" || payload.IdSRazlog == "" || payload.Vrepoc == "" || payload.IdTipob == "" || payload.ObId == "" || payload.IdSMrc == "" || payload.TipMan == "" {
		app.errorJSON(w, errors.New("mandatory data was not passed"))
		return
	} else {
		err := app.DB.InsertPiPiDDNIsklj(payload)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		resp = JSONResponse{
			Error:   false,
			Message: "Record inserted",
		}
	}
	//err = app.writeJSON(w, http.StatusOK, m)
	app.writeJSON(w, http.StatusAccepted, resp)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) insertPiPiDDNIskljP(w http.ResponseWriter, r *http.Request) {
	var payload models.PiPiDDNIsklj

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {

		app.errorJSON(w, err)
		return
	}

	var resp JSONResponse
	if payload.SynsoftId == "" || payload.DatSmene == "" || payload.IdSGrraz == "" || payload.IdSRazlog == "" || payload.Vrepoc == "" || payload.IdTipob == "" || payload.ObId == "" || payload.IdSMrc == "" || payload.TipMan == "" {
		app.errorJSON(w, errors.New("mandatory data was not passed"))
		return
	} else {
		err := app.DB.InsertPiPiDDNIskljP(payload)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		resp = JSONResponse{
			Error:   false,
			Message: "Record inserted",
		}
	}
	//err = app.writeJSON(w, http.StatusOK, m)
	app.writeJSON(w, http.StatusAccepted, resp)
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
	var resp JSONResponse
	if payload.SynsoftId == "" || payload.DatSmene == "" || payload.IdSGrraz == "" || payload.IdSRazlog == "" || payload.Vrepoc == "" || payload.IdTipob == "" || payload.ObId == "" || payload.IdSMrc == "" || payload.TipMan == "" {
		app.errorJSON(w, errors.New("mandatory data was not passed"))
		return
	} else {
		err := app.DB.UpdatePiPiDDNIsklj(payload)
		if err != nil {
			app.errorJSON(w, err)
			return
		}

		resp = JSONResponse{
			Error:   false,
			Message: "Record updated",
		}
	}
	app.writeJSON(w, http.StatusAccepted, resp)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) updatePiPiDDNIskljP(w http.ResponseWriter, r *http.Request) {
	var payload models.PiPiDDNIsklj

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	var resp JSONResponse
	if payload.SynsoftId == "" || payload.DatSmene == "" || payload.IdSGrraz == "" || payload.IdSRazlog == "" || payload.Vrepoc == "" || payload.IdTipob == "" || payload.ObId == "" || payload.IdSMrc == "" || payload.TipMan == "" {
		app.errorJSON(w, errors.New("mandatory data was not passed"))
		return
	} else {
		err := app.DB.UpdatePiPiDDNIskljP(payload)
		if err != nil {
			app.errorJSON(w, err)
			return
		}

		resp = JSONResponse{
			Error:   false,
			Message: "Record updated",
		}
	}
	app.writeJSON(w, http.StatusAccepted, resp)
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

	var resp JSONResponse
	if payload.SynsoftId == "" || payload.DatSmene == "" || payload.IdRadAPU == "" || payload.Id1SGruzr == "" || payload.Id1SUzrok == "" || payload.Vrepoc == "" || payload.IdSTipd == "" || payload.ObId == "" || payload.IdSMrc == "" || payload.IdSVrpd == "" {
		app.errorJSON(w, errors.New("mandatory data was not passed"))
		return
	} else {
		err = app.DB.InsertPiPiDDNIspad(payload)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		resp = JSONResponse{
			Error:   false,
			Message: "Record inserted",
		}
	}

	app.writeJSON(w, http.StatusAccepted, resp)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) insertPiPiDDNIspadP(w http.ResponseWriter, r *http.Request) {
	var payload models.PiPiDDNIspad

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {

		app.errorJSON(w, err)
		return
	}
	//fmt.Printf("Izgleda da je TipMan %s\n", payload.TipMan)

	var resp JSONResponse
	if payload.SynsoftId == "" || payload.DatSmene == "" || payload.IdRadAPU == "" || payload.Id1SGruzr == "" || payload.Id1SUzrok == "" || payload.Vrepoc == "" || payload.IdSTipd == "" || payload.ObId == "" || payload.IdSMrc == "" || payload.IdSVrpd == "" {
		app.errorJSON(w, errors.New("mandatory data was not passed"))
		return
	} else {
		err = app.DB.InsertPiPiDDNIspadP(payload)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		resp = JSONResponse{
			Error:   false,
			Message: "Record inserted",
		}
	}

	app.writeJSON(w, http.StatusAccepted, resp)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) updatePiPiDDNIspad(w http.ResponseWriter, r *http.Request) {
	var payload models.PiPiDDNIspad

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	var resp JSONResponse
	if payload.SynsoftId == "" || payload.DatSmene == "" || payload.IdRadAPU == "" || payload.Id1SGruzr == "" || payload.Id1SUzrok == "" || payload.Vrepoc == "" || payload.IdSTipd == "" || payload.ObId == "" || payload.IdSMrc == "" || payload.IdSVrpd == "" {
		app.errorJSON(w, errors.New("mandatory data was not passed"))
		return
	} else {
		err := app.DB.UpdatePiPiDDNIspad(payload)
		if err != nil {
			app.errorJSON(w, err)
			return
		}

		resp = JSONResponse{
			Error:   false,
			Message: "Record updated",
		}
	}
	app.writeJSON(w, http.StatusAccepted, resp)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) updatePiPiDDNIspadP(w http.ResponseWriter, r *http.Request) {
	var payload models.PiPiDDNIspad

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	var resp JSONResponse
	if payload.SynsoftId == "" || payload.DatSmene == "" || payload.IdRadAPU == "" || payload.Id1SGruzr == "" || payload.Id1SUzrok == "" || payload.Vrepoc == "" || payload.IdSTipd == "" || payload.ObId == "" || payload.IdSMrc == "" || payload.IdSVrpd == "" {
		app.errorJSON(w, errors.New("mandatory data was not passed"))
		return
	} else {
		err := app.DB.UpdatePiPiDDNIspadP(payload)
		if err != nil {
			app.errorJSON(w, err)
			return
		}

		resp = JSONResponse{
			Error:   false,
			Message: "Record updated",
		}
	}
	app.writeJSON(w, http.StatusAccepted, resp)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) deletePiPiDDN(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	/*id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}*/

	err := app.DB.DeletePiPiDDN(id)
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

func (app *application) deletePiPiDDNP(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	/*id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}*/

	err := app.DB.DeletePiPiDDNP(id)
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

/**** start nezavrseni_dog  *****/

func (app *application) getAllUnfinishedEventsNDC(w http.ResponseWriter, r *http.Request) {
	//year := chi.URLParam(r, "year")

	p, err := app.DB.GetAllUnfinishedEventsNDC()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, p)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllUnfinishedEventsNDCP(w http.ResponseWriter, r *http.Request) {
	//year := chi.URLParam(r, "year")

	p, err := app.DB.GetAllUnfinishedEventsNDCP()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, p)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getUnfinishedEventsByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	/*pipiddnID, err := strconv.Atoi(id)
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}*/

	p, err := app.DB.GetUnfinishedEventsByID(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, p)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getUnfinishedEventsByIDP(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	/*pipiddnID, err := strconv.Atoi(id)
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}*/

	p, err := app.DB.GetUnfinishedEventsByIDP(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, p)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) updateUnfinishedEvents(w http.ResponseWriter, r *http.Request) {
	var payload models.UnfinishedEventsUpdate

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	var resp JSONResponse
	if payload.DatSmene == "" || payload.TipSmene == "" || payload.SynsoftId == "" || payload.Id1SGruzr == "" || payload.Id1SUzrok == "" {
		app.errorJSON(w, errors.New("mandatory data was not passed"))
		return
	} else {
		err := app.DB.UpdateUnfinishedEvents(payload)
		if err != nil {
			app.errorJSON(w, err)
			return
		}

		resp = JSONResponse{
			Error:   false,
			Message: "Record updated",
		}
	}
	app.writeJSON(w, http.StatusAccepted, resp)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) updateUnfinishedEventsP(w http.ResponseWriter, r *http.Request) {
	var payload models.UnfinishedEventsUpdate

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	var resp JSONResponse
	if payload.DatSmene == "" || payload.TipSmene == "" || payload.SynsoftId == "" || payload.Id1SGruzr == "" || payload.Id1SUzrok == "" {
		app.errorJSON(w, errors.New("mandatory data was not passed"))
		return
	} else {
		err := app.DB.UpdateUnfinishedEventsP(payload)
		if err != nil {
			app.errorJSON(w, err)
			return
		}

		resp = JSONResponse{
			Error:   false,
			Message: "Record updated",
		}
	}
	app.writeJSON(w, http.StatusAccepted, resp)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

/**** end nezavrseni_dog  *****/

/***** start DDN_PREKID_ISP ****/
func (app *application) insertDDNInterruptionOfDelivery(w http.ResponseWriter, r *http.Request) {
	var payload models.DDNInterruptionOfDelivery

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {

		app.errorJSON(w, err)
		return
	}

	var resp JSONResponse
	if payload.IdSTipd == "" || payload.Vrepoc == "" || payload.IdTipob == "" || payload.ObId == "" || payload.IdSMrc == "" || payload.Ind == "" {
		app.errorJSON(w, errors.New("mandatory data was not passed"))
		return
	} else {
		err := app.DB.InsertDDNInterruptionOfDelivery(payload)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		resp = JSONResponse{
			Error:   false,
			Message: "Record inserted",
		}
	}
	//err = app.writeJSON(w, http.StatusOK, m)
	app.writeJSON(w, http.StatusAccepted, resp)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) insertDDNInterruptionOfDeliveryP(w http.ResponseWriter, r *http.Request) {
	var payload models.DDNInterruptionOfDelivery

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {

		app.errorJSON(w, err)
		return
	}

	var resp JSONResponse
	if payload.IdSTipd == "" || payload.Vrepoc == "" || payload.IdTipob == "" || payload.ObId == "" || payload.IdSMrc == "" || payload.Ind == "" {
		app.errorJSON(w, errors.New("mandatory data was not passed"))
		return
	} else {
		err := app.DB.InsertDDNInterruptionOfDeliveryP(payload)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		resp = JSONResponse{
			Error:   false,
			Message: "Record inserted",
		}
	}
	//err = app.writeJSON(w, http.StatusOK, m)
	app.writeJSON(w, http.StatusAccepted, resp)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) updateDDNInterruptionOfDelivery(w http.ResponseWriter, r *http.Request) {
	var payload models.DDNInterruptionOfDelivery

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	var resp JSONResponse
	if payload.IdSTipd == "" || payload.Vrepoc == "" || payload.IdTipob == "" || payload.ObId == "" || payload.IdSMrc == "" || payload.Ind == "" {
		app.errorJSON(w, errors.New("mandatory data was not passed"))
		return
	} else {
		err := app.DB.UpdateDDNInterruptionOfDelivery(payload)
		if err != nil {
			app.errorJSON(w, err)
			return
		}

		resp = JSONResponse{
			Error:   false,
			Message: "Record updated",
		}
	}
	app.writeJSON(w, http.StatusAccepted, resp)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) updateDDNInterruptionOfDeliveryP(w http.ResponseWriter, r *http.Request) {
	var payload models.DDNInterruptionOfDelivery

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	var resp JSONResponse
	if payload.IdSTipd == "" || payload.Vrepoc == "" || payload.IdTipob == "" || payload.ObId == "" || payload.IdSMrc == "" || payload.Ind == "" {
		app.errorJSON(w, errors.New("mandatory data was not passed"))
		return
	} else {
		err := app.DB.UpdateDDNInterruptionOfDeliveryP(payload)
		if err != nil {
			app.errorJSON(w, err)
			return
		}

		resp = JSONResponse{
			Error:   false,
			Message: "Record updated",
		}
	}
	app.writeJSON(w, http.StatusAccepted, resp)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) insertUpdateDDNInterruptionOfDelivery(w http.ResponseWriter, r *http.Request) {
	var payload models.DDNInterruptionOfDeliveryPayload

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	var resp JSONResponse
	if payload.SynsoftId == "" || payload.IdSTipd == "" || payload.Vrepoc == "" || payload.IdTipob == "" || payload.ObId == "" || payload.IdSMrc == "" || payload.P2TrafId == "" || payload.TipDogadjaja == "" {
		app.errorJSON(w, errors.New("mandatory data was not passed"))
		return
	} else {
		err := app.DB.InsertUpdateDDNInterruptionOfDelivery(payload)
		if err != nil {
			app.errorJSON(w, err)
			return
		}

		resp = JSONResponse{
			Error:   false,
			Message: "Record inserted/updated",
		}
	}
	app.writeJSON(w, http.StatusAccepted, resp)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) insertUpdateDDNInterruptionOfDeliveryP(w http.ResponseWriter, r *http.Request) {
	var payload models.DDNInterruptionOfDeliveryPayload

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	var resp JSONResponse
	if payload.SynsoftId == "" || payload.IdSTipd == "" || payload.Vrepoc == "" || payload.IdTipob == "" || payload.ObId == "" || payload.IdSMrc == "" || payload.P2TrafId == "" || payload.TipDogadjaja == "" {
		app.errorJSON(w, errors.New("mandatory data was not passed"))
		return
	} else {
		err := app.DB.InsertUpdateDDNInterruptionOfDeliveryP(payload)
		if err != nil {
			app.errorJSON(w, err)
			return
		}

		resp = JSONResponse{
			Error:   false,
			Message: "Record inserted/updated",
		}
	}
	app.writeJSON(w, http.StatusAccepted, resp)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

/**** definition of slice and sort  ****/
type Item struct {
	p    models.DDNInterruptionOfDeliveryPayload
	Date time.Time
}
type ByDate []Item

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].Date.Before(a[j].Date) }

/**** end definition of slice and sort  ****/

func (app *application) insertUpdateAllDDNInterruptionOfDelivery(w http.ResponseWriter, r *http.Request) {
	var payloads []models.DDNInterruptionOfDeliveryPayload
	dateFormat := "02.01.2006 15:04"
	err := app.readJSON(w, r, &payloads)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	var items []Item
	for _, payload := range payloads {
		date, err := time.Parse(dateFormat, payload.Vrepoc)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		item := Item{payload, date}
		items = append(items, item)
		/*fmt.Println(date)*/
	}
	sort.Sort(ByDate(items))
	for _, item := range items {
		/*fmt.Println(item.Date)*/
		if item.p.SynsoftId == "" || item.p.IdSTipd == "" || item.p.Vrepoc == "" || item.p.IdTipob == "" || item.p.ObId == "" || item.p.IdSMrc == "" || item.p.P2TrafId == "" || item.p.TipDogadjaja == "" {
			app.errorJSON(w, errors.New("mandatory data was not passed"))
			return
		} else {
			err := app.DB.InsertUpdateDDNInterruptionOfDelivery(item.p)
			if err != nil {
				app.errorJSON(w, err)
				return
			}
		}
	}
	/*for _, payload := range payloads {

		if payload.SynsoftId == "" || payload.IdSTipd == "" || payload.Vrepoc == "" || payload.IdTipob == "" || payload.ObId == "" || payload.IdSMrc == "" || payload.P2TrafId == "" || payload.TipDogadjaja == "" {
			app.errorJSON(w, errors.New("mandatory data was not passed"))
			return
		} else {
			err := app.DB.InsertUpdateDDNInterruptionOfDelivery(payload)
			if err != nil {
				app.errorJSON(w, err)
				return
			}
		}
	}*/

	//var resp JSONResponse
	var resp = JSONResponse{
		Error:   false,
		Message: "Record inserted/updated",
	}
	app.writeJSON(w, http.StatusAccepted, resp)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) deleteDDNInterruptionOfDelivery(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	/*id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}*/

	err := app.DB.DeleteDDNInterruptionOfDelivery(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponse{
		Error:   false,
		Message: "DDNInterruptionOfDelivery deleted",
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}

func (app *application) deleteDDNInterruptionOfDeliveryP(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	/*id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}*/

	err := app.DB.DeleteDDNInterruptionOfDeliveryP(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponse{
		Error:   false,
		Message: "DDNInterruptionOfDelivery deleted",
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}

func (app *application) getDDNInterruptionOfDeliveryNDC(w http.ResponseWriter, r *http.Request) {
	//year := chi.URLParam(r, "year")

	p, err := app.DB.GetDDNInterruptionOfDeliveryNDC()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, p)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getDDNInterruptionOfDeliveryNDCP(w http.ResponseWriter, r *http.Request) {
	//year := chi.URLParam(r, "year")

	p, err := app.DB.GetDDNInterruptionOfDeliveryNDCP()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, p)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getDDNInterruptionOfDeliveryByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	p, err := app.DB.GetDDNInterruptionOfDeliveryNDCByID(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, p)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getDDNInterruptionOfDeliveryByIDP(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	p, err := app.DB.GetDDNInterruptionOfDeliveryNDCByIDP(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, p)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

/*** start NOVITA ***/
func (app *application) getAllUnbalancedTrader(w http.ResponseWriter, r *http.Request) {
	//year := chi.URLParam(r, "year")

	p, err := app.DB.GetAllUnbalancedTrader()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, p)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

/*** end NOVITA ***/

/** start Check for PI_DOK **/

func (app *application) checkForPiDokP(w http.ResponseWriter, r *http.Request) {
	datSmene := chi.URLParam(r, "datsmene")

	idSMmrc, err := strconv.Atoi(chi.URLParam(r, "mrc"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	num, err := app.DB.CheckForPiDokYesterdayP(datSmene, idSMmrc)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	if num > 0 {
		app.errorJSON(w, fmt.Errorf("morate prvo zatvoriti Pogonski izvestaj za dan pre %s", datSmene))
		return
	}

	num, err = app.DB.CheckForPiDokTodayP(datSmene, idSMmrc)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	if num == 0 {
		app.errorJSON(w, fmt.Errorf("greska - Pogonski izvestaj za dan %s mora biti otvoren", datSmene))
		return
	}

	resp := JSONResponse{
		Error:   false,
		Message: "Regular status",
	}

	err = app.writeJSON(w, http.StatusOK, resp)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) closePgiP(w http.ResponseWriter, r *http.Request) {
	type Payload struct {
		Datsmene string `json:"datsmene"`
		Mrc      string `json:"mrc"`
	}

	var payload Payload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {

		app.errorJSON(w, err)
		return
	}

	idSMmrc, err := strconv.Atoi(payload.Mrc)
	var resp JSONResponse
	if payload.Datsmene == "" || payload.Mrc == "" {
		app.errorJSON(w, errors.New("mandatory data was not passed"))
		return
	} else {

		num, err := app.DB.CheckForPiDokYesterdayP(payload.Datsmene, idSMmrc)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		if num > 0 {
			app.errorJSON(w, fmt.Errorf("morate prvo zatvoriti Pogonski izvestaj za dan pre %s", payload.Datsmene))
			return
		}
		num, err = app.DB.CheckForPiDokTodayP(payload.Datsmene, idSMmrc)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		if num == 0 {
			app.errorJSON(w, fmt.Errorf("greska - Pogonski izvestaj za dan %s mora biti otvoren", payload.Datsmene))
			return
		}
		err = app.DB.ClosePgiP(payload.Datsmene, payload.Mrc)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		resp = JSONResponse{
			Error:   false,
			Message: "PGI is closed",
		}
	}

	app.writeJSON(w, http.StatusAccepted, resp)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) transferInPgiP(w http.ResponseWriter, r *http.Request) {
	type Payload struct {
		Datsmene string `json:"datsmene"`
		Mrc      string `json:"mrc"`
		Tip      string `json:"tip"`
	}

	var payload Payload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {

		app.errorJSON(w, err)
		return
	}

	var resp JSONResponse
	if payload.Datsmene == "" || payload.Mrc == "" || payload.Tip == "" {
		app.errorJSON(w, errors.New("mandatory data was not passed"))
		return
	} else {
		err := app.DB.TransferInPgiP(payload.Datsmene, payload.Mrc, payload.Tip)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		resp = JSONResponse{
			Error:   false,
			Message: "Transfer in PGI is complete",
		}
	}

	app.writeJSON(w, http.StatusAccepted, resp)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

/** end Check for PI_DOK **/

/****** end DDN_PREKID_ISP ****/

/**************** the another method ***********/
func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	// read json payload
	var requestPayload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	// validate user against database
	user, err := app.DB.GetUserByUsername(requestPayload.Username)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}

	// check password
	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}

	var roles []string
	for _, r := range user.UserRole {
		roles = append(roles, r.RoleCode)
	}

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

			/*if time.Unix(claims.ExpiresAt.Unix(), 0).Sub(time.Now()) > 30*time.Second {
				app.errorJSON(w, errors.New("refresh token does not need renewed yet"), http.StatusTooEarly)
				return
			}*/

			// get the user id from the token claims
			userID, err := strconv.Atoi(claims.Subject)
			if err != nil {
				app.errorJSON(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			user, err := app.DB.GetUserByID(userID)
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
