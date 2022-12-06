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
	/*mux.Get("", app.wrap(secure.ThenFunc(app.getDvDidf)))*/
	mux.Get("/signals/dv_didf", app.getDvDidf)
	mux.Get("/signals/tr_diff", app.getDiffTr)
	mux.Get("/signals/tr_dis_res", app.getDisTrRes)
	mux.Get("/signals/sp_dis_diff", app.getDisDiffSp)
	mux.Get("/signals/malfunction_in", app.getMalfunctionIn)
	mux.Get("/signals/apu", app.getAPU)
	mux.Get("/signals/dv_oc", app.getOCDV)
	mux.Get("/signals/tr12_oc", app.getOCTR12)
	mux.Get("/signals/trres_oc", app.getOCTRR)
	mux.Get("/signals/sp_oc_sc", app.getOCSP)
	mux.Get("/signals/dv_earthfault_oc", app.getEarthfaultOCDV)
	mux.Get("/signals/tr_earthfault_oc", app.getEarthfaultOCTR)
	mux.Get("/signals/sp_earthfault_oc", app.getEarthfaultOCSP)
	mux.Get("/signals/dir_earthfault_oc", app.getDirEarthfaultOC)
	mux.Get("/signals/tp_send_rcvd", app.getTPSendRcdv)
	mux.Get("/signals/circuitbreaker", app.getCircuitbreaker)
	mux.Get("/signals/bbp_bf_trip", app.getBBPBFtrip)
	mux.Get("/signals/non_electrical", app.getNonElectrical)
	mux.Get("/signals/bbp_bb_trip", app.getBBPBBtrip)
	mux.Get("/signals/bf_trip", app.getBFtrip)
	mux.Get("/groups_of_causes", app.getGroupsOfCauses)
	mux.Get("/causes", app.getCauses)
	mux.Get("/groups_of_reasons", app.getGroupOfReasons)
	mux.Get("/reasons", app.getReasons)
	mux.Get("/weather_conditions", app.getWeatherConditions)
	mux.Get("/categories_od_events", app.getCategoriesOfEvents)
	mux.Get("/overhead_lines", app.getOHL)
	mux.Get("/power_cables", app.getPowerCables)
	mux.Get("/substations", app.getSubstations)
	mux.Get("/feeders", app.getFeeders)
	mux.Get("/protection_devices", app.getProtectionDevices)
	mux.Get("/power_transformers", app.getPowerTransformers)
	mux.Get("/disconnectors", app.getDisconnectors)

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(app.authRequired)

		mux.Get("/signals/dv_didf", app.getDvDidf)
	})

	return mux

}
