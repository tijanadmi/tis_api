package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

/*func (app *application) wrap(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}*/

func (app *application) routes() http.Handler {
	/*router := httprouter.New()
	secure := alice.New(app.checkToken)

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	router.HandlerFunc(http.MethodPost, "/signin", app.Signin)
	router.HandlerFunc(http.MethodGet, "/signals/dv_didf", app.getDvDidf)
	router.HandlerFunc(http.MethodGet, "/signals/tr_diff", app.getDiffTr)
	router.HandlerFunc(http.MethodGet, "/signals/tr_dis_res", app.getDisTrRes)
	router.HandlerFunc(http.MethodGet, "/signals/sp_dis_diff", app.getDisDiffSp)
	router.HandlerFunc(http.MethodGet, "/signals/malfunction_in", app.getMalfunctionIn)
	router.HandlerFunc(http.MethodGet, "/signals/apu", app.getAPU)
	router.HandlerFunc(http.MethodGet, "/signals/dv_oc", app.getOCDV)

	router.POST("/dv_didf", app.wrap(secure.ThenFunc(app.getDvDidf)))
	return app.enableCORS(router)*/
	// create a router mux
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	mux.Post("/authenticate", app.authenticate)
	mux.Get("/refresh", app.refreshToken)
	mux.Get("/logout", app.logout)

	mux.Get("/status", app.statusHandler)
	mux.Post("/signin", app.Signin)

	mux.Put("/pipiddn/operation/0", app.insertPiPiDDNIsklj)
	mux.Put("/pipiddn/outage/0", app.insertPiPiDDNIspad)

	mux.Patch("/pipiddn/operation/{id}", app.updatePiPiDDNIsklj)
	mux.Patch("/pipiddn/outage/{id}", app.updatePiPiDDNIspad)

	mux.Delete("/pipiddn/{id}", app.deletePiPiDDN)

	mux.Route("/dwh", func(mux chi.Router) {
		mux.Use(app.authRequiredDWH)
		mux.Get("/weather/{year}", app.getWeather)
		mux.Get("/weatherforecast", app.getWeatherForecast)
		mux.Get("/weatherhistory/{year}", app.getWeatherHistory)

		mux.Get("/dozvole1/{year}", app.getPermissions1)
		mux.Get("/dozvole23/{year}", app.getPermissions23)
		mux.Get("/zahtevi1/{year}", app.getRequests1)
		mux.Get("/zahtevi23/{year}", app.getRequests23)
		mux.Get("/ispadi/{year}", app.getOutages)
		mux.Get("/iskljucenja/{year}", app.getExclusions)
		mux.Get("/planovi/{year}", app.getPlans)
	})

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(app.authRequiredNDC)

		mux.Get("/signals/{id}", app.getOneSignal)
		mux.Get("/signals/dvdidf", app.getDvDidf)
		mux.Get("/signals/trdiff", app.getDiffTr)
		mux.Get("/signals/trdisres", app.getDisTrRes)
		mux.Get("/signals/spdisdiff", app.getDisDiffSp)
		mux.Get("/signals/malfunctionin", app.getMalfunctionIn)
		mux.Get("/signals/malfunctionin/{id}", app.getOneMalfunctionIn)
		mux.Get("/signals/apu", app.getAPU)
		mux.Get("/signals/apu/{id}", app.getOneAPU)
		mux.Get("/signals/dvoc", app.getOCDV)
		mux.Get("/signals/tr12oc", app.getOCTR12)
		mux.Get("/signals/trresoc", app.getOCTRR)
		mux.Get("/signals/spocsc", app.getOCSP)
		mux.Get("/signals/dvearthfaultoc", app.getEarthfaultOCDV)
		mux.Get("/signals/trearthfaultoc", app.getEarthfaultOCTR)
		mux.Get("/signals/spearthfaultoc", app.getEarthfaultOCSP)
		mux.Get("/signals/direarthfaultoc", app.getDirEarthfaultOC)
		mux.Get("/signals/tpsendrcvd", app.getTPSendRcdv)
		mux.Get("/signals/circuitbreaker", app.getCircuitbreaker)
		mux.Get("/signals/bbpbftrip", app.getBBPBFtrip)
		mux.Get("/signals/nonelectrical", app.getNonElectrical)
		mux.Get("/signals/bbpbbtrip", app.getBBPBBtrip)
		mux.Get("/signals/bftrip", app.getBFtrip)
		mux.Get("/groupsofcauses", app.getGroupsOfCauses)
		mux.Get("/groupsofcauses/{id}", app.getOneGroupOfCauses)
		mux.Get("/causes", app.getCauses)
		mux.Get("/causes/{id}", app.getOneCause)
		mux.Get("/groupsofreasons", app.getGroupOfReasons)
		mux.Get("/groupsofreasons/{id}", app.getOneGroupOfReason)
		mux.Get("/reasons", app.getReasons)
		mux.Get("/reasons/{id}", app.getOneReason)
		mux.Get("/weatherconditions", app.getWeatherConditions)
		mux.Get("/weatherconditions/{id}", app.getOneWeatherCondition)
		mux.Get("/categoriesofevents", app.getCategoriesOfEvents)
		mux.Get("/categoriesofevents/{id}", app.getOneCategoryOfEvents)
		mux.Get("/overheadlines", app.getOHL)
		mux.Get("/powercables", app.getPowerCables)
		mux.Get("/substations", app.getSubstations)
		mux.Get("/feeders", app.getFeeders)
		mux.Get("/protectiondevices", app.getProtectionDevices)
		mux.Get("/powertransformers", app.getPowerTransformers)
		mux.Get("/disconnectors", app.getDisconnectors)
		mux.Get("/workpermissions", app.getWorkPermissions)
		mux.Get("/workineenetwork", app.getWorkInEENetwork)
	})

	return mux

}
